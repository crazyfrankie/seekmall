package ioc

import (
	"github.com/crazyfrankie/seekmall/rpc_gen/cart"
	"github.com/crazyfrankie/seekmall/rpc_gen/product"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/crazyfrankie/seekmall/rpc_gen/sm"
	"github.com/crazyfrankie/seekmall/rpc_gen/user"
)

func InitUserClient(cli *clientv3.Client) user.UserServiceClient {
	builder, err := resolver.NewBuilder(cli)
	conn, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(builder),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return user.NewUserServiceClient(conn)
}

func InitSmsClient(cli *clientv3.Client) sm.SmsServiceClient {
	builder, err := resolver.NewBuilder(cli)
	conn, err := grpc.Dial("etcd:///service/sms",
		grpc.WithResolvers(builder),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return sm.NewSmsServiceClient(conn)
}

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

func InitCartClient(cli *clientv3.Client) cart.CartServiceClient {
	builder, err := resolver.NewBuilder(cli)
	conn, err := grpc.Dial("etcd:///service/cart",
		grpc.WithResolvers(builder),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return cart.NewCartServiceClient(conn)
}
