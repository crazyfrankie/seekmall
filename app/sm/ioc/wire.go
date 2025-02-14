//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/seekmall/app/sm/biz/repository"
	"github.com/crazyfrankie/seekmall/app/sm/biz/repository/cache"
	"github.com/crazyfrankie/seekmall/app/sm/biz/service"
	"github.com/crazyfrankie/seekmall/app/sm/biz/service/sms/memory"
	"github.com/crazyfrankie/seekmall/app/sm/config"
	"github.com/crazyfrankie/seekmall/app/sm/rpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func InitCache() redis.Cmdable {
	cli := redis.NewClient(&redis.Options{
		Addr: config.GetConf().Redis.Addr,
	})

	return cli
}

func InitRegistry() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{config.GetConf().ETCD.Addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	return cli
}

func InitServer() *rpc.Server {
	wire.Build(
		InitCache,
		cache.NewSmsCache,
		repository.NewSmsRepo,
		memory.NewMemorySms,
		service.NewSmsServer,
		InitRegistry,
		rpc.NewServer,
	)

	return new(rpc.Server)
}
