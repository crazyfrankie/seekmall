//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/seekmall/app/ms"
	"github.com/crazyfrankie/seekmall/app/product"
	"github.com/crazyfrankie/seekmall/app/user"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type App struct {
	Server *gin.Engine
}

var BaseSet = wire.NewSet(InitDB, InitCache)

func InitApp() *App {
	wire.Build(
		BaseSet,
		ms.InitModule,
		user.InitModule,
		product.InitModule,
		InitMws,
		InitWeb,

		wire.FieldsOf(new(*user.Module), "Hdl"),
		wire.FieldsOf(new(*product.Module), "Hdl"),
		wire.Struct(new(App), "*"),
	)

	return new(App)
}
