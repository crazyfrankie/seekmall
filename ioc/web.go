package ioc

import (
	"github.com/crazyfrankie/seekmall/app/pkg/mws"
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/seekmall/app/product"
	"github.com/crazyfrankie/seekmall/app/user"
)

func InitWeb(mws []gin.HandlerFunc, user *user.Handler, product *product.Handler) *gin.Engine {
	server := gin.Default()

	server.Use(mws...)
	user.RegisterRoute(server)
	product.RegisterRoute(server)

	return server
}

func InitMws() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		mws.NewAuthBuilder().
			IgnorePath("/api/user/send-code").
			IgnorePath("/api/user/verify-code").
			IgnorePath("/api/user/login").
			Auth(),
	}
}
