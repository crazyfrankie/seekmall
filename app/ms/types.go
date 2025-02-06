package ms

import "github.com/crazyfrankie/seekmall/app/ms/service"

type SmsSvc = service.SmsSvc

type Module struct {
	Sms SmsSvc
}
