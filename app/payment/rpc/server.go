package rpc

import (
	"net"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"

	"github.com/crazyfrankie/seekmall/app/payment/biz/service"
	"github.com/crazyfrankie/seekmall/app/payment/config"
	"github.com/crazyfrankie/seekmall/rpc_gen/payment"
)

type Server struct {
	*grpc.Server
	Addr   string
	client *clientv3.Client
}

func NewServer(p *service.PaymentServer, cli *clientv3.Client) *Server {
	s := grpc.NewServer()

	payment.RegisterPaymentServiceServer(s, p)

	return &Server{
		Server: s,
		Addr:   config.GetConf().Server.Addr,
		client: cli,
	}
}

func (s *Server) Serve() error {
	conn, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	return s.Server.Serve(conn)
}
