package rpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"

	"github.com/crazyfrankie/seekmall/app/cart/biz/service"
	"github.com/crazyfrankie/seekmall/app/cart/config"
)

type Server struct {
	*grpc.Server
	Addr   string
	client *clientv3.Client
}

func NewServer(cart *service.CartServer, client *clientv3.Client) *Server {
	s := grpc.NewServer()
	cart.RegisterServer(s)

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

	err = registerService(s.client, s.Addr)
	if err != nil {
		return err
	}

	return s.Server.Serve(conn)
}

func registerService(cli *clientv3.Client, port string) error {
	em, err := endpoints.NewManager(cli, "service/cart")
	if err != nil {
		return err
	}

	addr := "127.0.0.1" + port
	serviceKey := "/service/cart/" + addr

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	leaseResp, err := cli.Grant(ctx, 60)
	if err != nil {
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
					log.Println("KeepAlive channel closed")
					return
				}
				fmt.Println("Lease renewed")
			case <-ctx.Done():
				return
			}
		}
	}()

	return err
}
