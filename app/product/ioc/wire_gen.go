// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package ioc

import (
	"fmt"
	"github.com/crazyfrankie/seekmall/app/product/biz/repository"
	"github.com/crazyfrankie/seekmall/app/product/biz/repository/dao"
	"github.com/crazyfrankie/seekmall/app/product/biz/service"
	"github.com/crazyfrankie/seekmall/app/product/config"
	"github.com/crazyfrankie/seekmall/app/product/rpc"
	"go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

// Injectors from wire.go:

func InitServer() *rpc.Server {
	db := InitDB()
	sellerDao := dao.NewSellerDao(db)
	purchaserDao := dao.NewPurchaserDao(db)
	productRepo := repository.NewProductRepo(sellerDao, purchaserDao)
	productServer := service.NewProductServer(productRepo)
	client := InitRegistry()
	server := rpc.NewServer(productServer, client)
	return server
}

// wire.go:

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(config.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

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
