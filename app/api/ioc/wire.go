//go:build wireinject

package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/crazyfrankie/seekmall/app/api/handler"
	"github.com/crazyfrankie/seekmall/app/api/pkg/mws"
)

func InitMws() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		mws.NewAuthBuilder().
			IgnorePath("/api/user/send-code").
			IgnorePath("/api/user/verify-code").
			Auth(),
	}
}

func InitWeb(mws []gin.HandlerFunc, user *handler.UserHandler, product *handler.ProductHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mws...)

	user.RegisterRoute(server)
	product.RegisterRoute(server)

	return server
}

func InitGin() *gin.Engine {
	wire.Build(
		InitUserClient,
		InitSmsClient,
		InitProductClient,
		handler.NewUserHandler,
		handler.NewProductHandler,
		InitMws,
		InitWeb,
	)
	return new(gin.Engine)
}
