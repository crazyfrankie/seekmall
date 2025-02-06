//go:build wireinject

package ms

import (
	"github.com/crazyfrankie/seekmall/app/ms/repository"
	"github.com/crazyfrankie/seekmall/app/ms/repository/cache"
	"github.com/crazyfrankie/seekmall/app/ms/service"
	"github.com/crazyfrankie/seekmall/app/ms/service/sms/memory"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func InitModule(cmd redis.Cmdable) *Module {
	wire.Build(
		cache.NewSmsCache,
		repository.NewSmsRepo,
		memory.NewMemorySms,
		service.NewSmsService,

		wire.Struct(new(Module), "*"),
	)
	return new(Module)
}
