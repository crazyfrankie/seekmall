package repository

import (
	"context"

	"github.com/crazyfrankie/seekmall/app/ms/repository/cache"
)

type SmsRepo struct {
	cache *cache.SmsCache
}

func NewSmsRepo(cache *cache.SmsCache) *SmsRepo {
	return &SmsRepo{cache: cache}
}

func (r *SmsRepo) Store(ctx context.Context, receiver, code string) error {
	return r.cache.Store(ctx, receiver, code)
}

func (r *SmsRepo) Verify(ctx context.Context, receiver, code string) error {
	return r.cache.Verify(ctx, receiver, code)
}
