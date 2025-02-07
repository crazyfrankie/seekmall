package rpc

import (
	"net"
	
	"google.golang.org/grpc"

	"github.com/crazyfrankie/seekmall/app/product/biz/service"
	"github.com/crazyfrankie/seekmall/app/product/config"
)

type Server struct {
	*grpc.Server
	Addr string
}

func NewServer(product *service.ProductServer) *Server {
	s := grpc.NewServer()
	product.RegisterServer(s)

	return &Server{
		Server: s,
		Addr:   config.GetConf().Server.Addr,
	}
}

func (s *Server) Serve() error {
	conn, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}

	return s.Server.Serve(conn)
}
