package rpc

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	"github.com/crazyfrankie/seekmall/app/user/biz/service"
	"github.com/crazyfrankie/seekmall/app/user/config"
)

type Server struct {
	*grpc.Server
	Addr   string
	client *clientv3.Client
}

func NewServer(u *service.UserServer, client *clientv3.Client) *Server {
	s := grpc.NewServer()
	u.RegisterServer(s)

	return &Server{
		Server: s,
		Addr:   config.GetConf().Server.Addr,
		client: client,
	}
}

func (s *Server) Serve() error {
	conn, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	err = serviceRegister(s.client, s.Addr)
	if err != nil {
		return err
	}

	return s.Server.Serve(conn)
}

func serviceRegister(cli *clientv3.Client, port string) error {
	em, err := endpoints.NewManager(cli, "service/user")
	if err != nil {
		return err
	}

	addr := "127.0.0.1" + port
	serviceKey := "/services/user/" + addr
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	leaseResp, err := cli.Grant(context.Background(), 5)
	if err != nil {
		log.Fatalf("failed to grant lease: %v", err)
		return err
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = em.AddEndpoint(ctx, serviceKey, endpoints.Endpoint{Addr: addr}, clientv3.WithLease(leaseResp.ID))

	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		ch, err := cli.KeepAlive(ctx, leaseResp.ID)
		if err != nil {
			log.Fatalf("KeepAlive failed: %v", err)
		}

		for {
			select {
			case _, ok := <-ch:
				if !ok {
					fmt.Println("KeepAlive channel closed")
				}
				fmt.Println("Lease renewed")
			case <-ctx.Done():
				return
			}
		}
	}()

	return err
}
