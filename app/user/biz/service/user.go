package service

import (
	"context"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/seekmall/app/user/biz/repository"
	"github.com/crazyfrankie/seekmall/app/user/biz/repository/dao"
	"github.com/crazyfrankie/seekmall/app/user/pkg/constants"
	"github.com/crazyfrankie/seekmall/app/user/pkg/mws"
	"github.com/crazyfrankie/seekmall/rpc_gen/sm"
	"github.com/crazyfrankie/seekmall/rpc_gen/user"
)

type UserServer struct {
	repo      *repository.UserRepo
	smsClient sm.SmsServiceClient
	user.UnimplementedUserServiceServer
}

func NewUserServer(repo *repository.UserRepo, smsClient sm.SmsServiceClient) *UserServer {
	return &UserServer{repo: repo, smsClient: smsClient}
}

func (s *UserServer) RegisterServer(server *grpc.Server) {
	user.RegisterUserServiceServer(server, s)
}

func (s *UserServer) SendCode(ctx context.Context, req *user.SendCodeRequest) (*user.SendCodeResponse, error) {
	var biz string
	phone := req.GetPhone()
	u, err := s.repo.FindByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	if u.Id == 0 {
		biz = constants.Register
	} else {
		biz = constants.Login
	}

	_, err = s.smsClient.SendSms(ctx, &sm.SendSmsRequest{
		Biz:   biz,
		Phone: phone,
	})
	if err != nil {
		return nil, err
	}

	return &user.SendCodeResponse{
		Biz: biz,
	}, nil
}

func (s *UserServer) VerifyCode(ctx context.Context, req *user.VerifyCodeRequest) (*user.VerifyCodeResponse, error) {
	biz, phone, code := req.GetBiz(), req.GetPhone(), req.GetCode()
	_, err := s.smsClient.VerifySms(ctx, &sm.VerifySmsRequest{
		Biz:   biz,
		Phone: phone,
		Code:  code,
	})
	if err != nil {
		return nil, err
	}

	var uid int
	if biz == constants.Register {
		u := &dao.User{
			Phone: phone,
			Name:  phone,
		}
		err := s.repo.Create(ctx, u)
		if err != nil {
			return nil, err
		}
		uid = u.Id
	}

	var token string
	token, err = mws.GenerateToken(uid)
	if err != nil {
		return nil, err
	}

	return &user.VerifyCodeResponse{Token: token}, nil
}

func (s *UserServer) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	u, err := s.repo.FindById(ctx, int(req.GetUid()))
	if err != nil {
		return nil, err
	}

	return &user.GetUserInfoResponse{
		User: &user.User{
			Id:    int32(u.Id),
			Name:  u.Name,
			Phone: u.Phone,
		},
	}, nil
}
