//go:build wireinject

package ioc

import (
	"fmt"
	"os"

	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/crazyfrankie/seekmall/app/user/biz/repository"
	"github.com/crazyfrankie/seekmall/app/user/biz/repository/dao"
	"github.com/crazyfrankie/seekmall/app/user/biz/rpc"
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

type App struct {
	server *rpc.Server
}

func InitApp() *App {
	wire.Build(
		InitDB,
		dao.NewUserDao,
		repository.NewUserRepo,
		service.NewUserServer,
		rpc.InitSmsClient,
		rpc.NewServer,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
