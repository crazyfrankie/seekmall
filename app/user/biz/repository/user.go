package repository

import (
	"context"

	"github.com/crazyfrankie/seekmall/app/user/biz/repository/dao"
)

type UserRepo struct {
	dao *dao.UserDao
}

func NewUserRepo(d *dao.UserDao) *UserRepo {
	return &UserRepo{dao: d}
}

func (r *UserRepo) Create(ctx context.Context, u *dao.User) error {
	return r.dao.Create(ctx, u)
}

func (r *UserRepo) FindByPhone(ctx context.Context, phone string) (dao.User, error) {
	return r.dao.FindByPhone(ctx, phone)
}

func (r *UserRepo) FindById(ctx context.Context, uid int) (dao.User, error) {
	return r.dao.FindById(ctx, uid)
}
