package service

import (
	"context"

	"github.com/crazyfrankie/seekmall/app/pkg/mws"
	"github.com/crazyfrankie/seekmall/app/user/repository"
	"github.com/crazyfrankie/seekmall/app/user/repository/dao"
)

type UserSvc interface {
	FindOrCreateUser(ctx context.Context, phone string) (string, error)
	FindById(ctx context.Context, uid int) (dao.User, error)
}

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) UserSvc {
	return &UserService{repo: repo}
}

func (s *UserService) FindOrCreateUser(ctx context.Context, phone string) (string, error) {
	u, err := s.repo.FindByPhone(ctx, phone)
	if err != nil {
		return "", err
	}
	var uid int
	if u.Id == 0 {
		newUser := &dao.User{
			Phone: phone,
			Name:  phone,
		}
		err := s.repo.Create(ctx, newUser)
		if err != nil {
			return "", err
		}
		uid = newUser.Id
	} else {
		uid = u.Id
	}

	var token string
	token, err = mws.GenerateToken(uid)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) FindById(ctx context.Context, uid int) (dao.User, error) {
	return s.repo.FindById(ctx, uid)
}
