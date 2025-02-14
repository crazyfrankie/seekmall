package repository

import (
	"context"

	"github.com/crazyfrankie/seekmall/app/cart/biz/repository/dao"
)

type CartRepo struct {
	dao *dao.CartDao
}

func NewCartRepo(d *dao.CartDao) *CartRepo {
	return &CartRepo{dao: d}
}

func (r *CartRepo) CreateItem(ctx context.Context, i *dao.Item) error {
	return r.dao.CreateItem(ctx, i)
}

func (r *CartRepo) FindByUserId(ctx context.Context, uid int32) ([]*dao.Item, error) {
	return r.dao.FindByUserId(ctx, uid)
}

func (r *CartRepo) DeleteCartsById(ctx context.Context, uid int32) error {
	return r.dao.DeleteCartsById(ctx, uid)
}
