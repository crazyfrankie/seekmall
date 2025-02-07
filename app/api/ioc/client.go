package ioc

import (
	"github.com/crazyfrankie/seekmall/rpc_gen/product"
	"google.golang.org/grpc"

	"github.com/crazyfrankie/seekmall/rpc_gen/sm"
	"github.com/crazyfrankie/seekmall/rpc_gen/user"
)

func InitUserClient() user.UserServiceClient {
	conn, err := grpc.NewClient("localhost:8081")
	if err != nil {
		panic(err)
	}

	return user.NewUserServiceClient(conn)
}

func InitSmsClient() sm.SmsServiceClient {
	conn, err := grpc.NewClient("localhost:8082")
	if err != nil {
		panic(err)
	}

	return sm.NewSmsServiceClient(conn)
}

func InitProductClient() product.ProductServiceClient {
	conn, err := grpc.NewClient("localhost:8083")
	if err != nil {
		panic(err)
	}

	return product.NewProductServiceClient(conn)
}
