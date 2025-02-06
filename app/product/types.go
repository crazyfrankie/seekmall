package product

import "github.com/crazyfrankie/seekmall/app/product/handler"

type Handler = handler.ProductHandler

type Module struct {
	Hdl *Handler
}
