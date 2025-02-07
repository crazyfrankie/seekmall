package repository

import (
	"context"
	"github.com/crazyfrankie/seekmall/app/sm/biz/repository/cache"
)

type SmsRepo struct {
	cache *cache.SmsCache
}

func NewSmsRepo(cache *cache.SmsCache) *SmsRepo {
	return &SmsRepo{cache: cache}
}

func (r *SmsRepo) Store(ctx context.Context, biz, receiver, code string) error {
	return r.cache.Store(ctx, biz, receiver, code)
}

func (r *SmsRepo) Verify(ctx context.Context, biz, receiver, code string) error {
	return r.cache.Verify(ctx, biz, receiver, code)
}
