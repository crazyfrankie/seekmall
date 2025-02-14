//go:build wireinject

package ioc

import (
	"fmt"
	rpc2 "github.com/crazyfrankie/seekmall/app/user/rpc"
	"os"
	"time"

	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/crazyfrankie/seekmall/app/user/biz/repository"
	"github.com/crazyfrankie/seekmall/app/user/biz/repository/dao"
	"github.com/crazyfrankie/seekmall/app/user/biz/service"
	"github.com/crazyfrankie/seekmall/app/user/config"
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

	db.AutoMigrate(&dao.User{})

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

func InitServer() *rpc2.Server {
	wire.Build(
		InitDB,
		InitRegistry,
		dao.NewUserDao,
		repository.NewUserRepo,
		service.NewUserServer,
		rpc2.InitSmsClient,
		rpc2.NewServer,
	)
	return new(rpc2.Server)
}
