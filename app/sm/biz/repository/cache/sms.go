package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/crazyfrankie/seekmall/app/sm/pkg/constants"
)

var (
	//go:embed lua/set_code.lua
	luaSetCode string
	//go:embed lua/verify_code.lua
	luaVerifyCode string
)

type SmsCache struct {
	cmd redis.Cmdable
}

func NewSmsCache(cmd redis.Cmdable) *SmsCache {
	return &SmsCache{cmd: cmd}
}

func (c *SmsCache) Store(ctx context.Context, biz, receiver string, code string) error {
	res, err := c.cmd.Eval(ctx, luaSetCode, []string{c.key(biz, receiver)}, code).Result()

	if err != nil {
		return err
	}
	switch res {
	case 0:
		return nil
	case -1:
		return constants.VerifyTooMany
	}

	return errors.New("internal server error")
}

func (c *SmsCache) Verify(ctx context.Context, biz, receiver string, code string) error {
	res, err := c.cmd.Eval(ctx, luaVerifyCode, []string{c.key(biz, receiver)}, code).Int()
	if err != nil {
		return err
	}

	switch res {
	case 0:
		// 毫无问题
		return nil
	case -1:
		// 发送太频繁
		return constants.SendTooMany
	}

	return errors.New("internal server error")
}

func (c *SmsCache) key(biz, receiver string) string {
	return fmt.Sprintf("phone_code:%s:%s", biz, receiver)
}
