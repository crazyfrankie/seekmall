package memory

import (
	"context"
	"fmt"
	"github.com/crazyfrankie/seekmall/app/ms/service/sms"
)

type MemorySms struct {
}

func NewMemorySms() sms.Service {
	return &MemorySms{}
}

func (m *MemorySms) Send(ctx context.Context, biz string, args []string, numbers ...string) error {
	fmt.Println(args)
	return nil
}
