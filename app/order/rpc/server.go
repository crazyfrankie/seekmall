package rpc

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	Addr   string
	client *clientv3.Client
}

func NewServer(client *clientv3.Client) *Server {
	s := grpc.NewServer()

	return &Server{
		Server: s,
		Addr:   "",
		client: client,
	}
}
