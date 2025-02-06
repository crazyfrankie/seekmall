package user

import "github.com/crazyfrankie/seekmall/app/user/handler"

type Handler = handler.UserHandler
type Module struct {
	Hdl *Handler
}
