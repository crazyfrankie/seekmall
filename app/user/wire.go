//go:build wireinject

package user

import (
	"github.com/crazyfrankie/seekmall/app/ms"
	"github.com/crazyfrankie/seekmall/app/user/handler"
	"github.com/crazyfrankie/seekmall/app/user/repository"
	"github.com/crazyfrankie/seekmall/app/user/repository/dao"
	"github.com/crazyfrankie/seekmall/app/user/service"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitModule(db *gorm.DB, sms *ms.Module) *Module {
	wire.Build(
		dao.NewUserDao,
		repository.NewUserRepo,
		service.NewUserService,
		handler.NewUserHandler,

		wire.FieldsOf(new(*ms.Module), "Sms"),
		wire.Struct(new(Module), "*"),
	)
	return new(Module)
}
