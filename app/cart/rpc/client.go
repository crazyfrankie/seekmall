package rpc

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/crazyfrankie/seekmall/rpc_gen/product"
)

func InitProductClient(cli *clientv3.Client) product.ProductServiceClient {
	builder, err := resolver.NewBuilder(cli)

	conn, err := grpc.Dial("etcd:///service/product",
		grpc.WithResolvers(builder),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return product.NewProductServiceClient(conn)
}
