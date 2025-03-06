//go:build wireinject

package ioc

import (
	"github.com/crazyfrankie/seekmall/app/api/config"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"

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

func InitWeb(mws []gin.HandlerFunc, user *handler.UserHandler,
	product *handler.ProductHandler, cart *handler.CartHandler, payment *handler.PaymentHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mws...)

	user.RegisterRoute(server)
	product.RegisterRoute(server)
	cart.RegisterRoute(server)
	payment.RegisterRoute(server)

	return server
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

func InitNotify() *notify.Handler {
	return notify.NewEmptyHandler()
}

func InitGin() *gin.Engine {
	wire.Build(
		InitRegistry,
		InitNotify,
		InitUserClient,
		InitSmsClient,
		InitProductClient,
		InitCartClient,
		InitPaymentClient,
		handler.NewUserHandler,
		handler.NewProductHandler,
		handler.NewCartHandler,
		handler.NewPaymentHandler,
		InitMws,
		InitWeb,
	)
	return new(gin.Engine)
}
