//go:build wireinject

package ioc

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/wire"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/crazyfrankie/seekmall/app/payment/biz/repository"
	"github.com/crazyfrankie/seekmall/app/payment/biz/repository/dao"
	"github.com/crazyfrankie/seekmall/app/payment/biz/service"
	"github.com/crazyfrankie/seekmall/app/payment/biz/service/wechat"
	"github.com/crazyfrankie/seekmall/app/payment/config"
	"github.com/crazyfrankie/seekmall/app/payment/rpc"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(config.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&dao.Payment{})

	return db
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

func InitNativeService() *native.NativeApiService {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cli, err := core.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	return &native.NativeApiService{Client: cli}
}

func InitServer() *rpc.Server {
	wire.Build(
		InitDB,
		InitRegistry,
		dao.NewPaymentDao,
		repository.NewPaymentRepo,
		InitNativeService,
		wechat.NewNativePayService,
		service.NewPaymentServer,
		rpc.NewServer,
	)

	return new(rpc.Server)
}
