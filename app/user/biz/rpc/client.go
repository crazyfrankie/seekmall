package rpc

import (
	"google.golang.org/grpc"
	
	"github.com/crazyfrankie/seekmall/rpc_gen/sm"
)

func InitSmsClient() sm.SmsServiceClient {
	conn, err := grpc.NewClient("localhost:8082")
	if err != nil {
		panic(err)
	}

	return sm.NewSmsServiceClient(conn)
}
