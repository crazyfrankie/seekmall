package ioc

import (
	"github.com/redis/go-redis/v9"

	"github.com/crazyfrankie/seekmall/config"
)

func InitCache() redis.Cmdable {
	return redis.NewClient(&redis.Options{
		Addr: config.GetConf().Redis.Addr,
	})
}
