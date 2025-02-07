//go:build wireinject

package ioc

import (
	"fmt"
	"github.com/google/wire"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/crazyfrankie/seekmall/app/product/biz/repository"
	"github.com/crazyfrankie/seekmall/app/product/biz/repository/dao"
	"github.com/crazyfrankie/seekmall/app/product/biz/rpc"
	"github.com/crazyfrankie/seekmall/app/product/biz/service"
	"github.com/crazyfrankie/seekmall/app/product/config"
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

	return db
}

type App struct {
	server *rpc.Server
}

func InitApp() *App {
	wire.Build(
		InitDB,
		dao.NewPurchaserDao,
		dao.NewSellerDao,
		repository.NewProductRepo,
		service.NewProductServer,
		rpc.NewServer,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
