//go:build wireinject

package product

import (
	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/crazyfrankie/seekmall/app/product/handler"
	"github.com/crazyfrankie/seekmall/app/product/repository"
	"github.com/crazyfrankie/seekmall/app/product/repository/dao"
	"github.com/crazyfrankie/seekmall/app/product/service"
)

func InitModule(db *gorm.DB) *Module {
	wire.Build(
		dao.NewSellerDao,
		dao.NewPurchaserDao,
		repository.NewProductRepo,
		service.NewProductService,
		handler.NewProductHandler,

		wire.Struct(new(Module), "*"),
	)
	return new(Module)
}
