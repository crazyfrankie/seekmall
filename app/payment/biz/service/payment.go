package service

import (
	"context"
	"database/sql"

	"github.com/crazyfrankie/seekmall/app/payment/biz/repository"
	"github.com/crazyfrankie/seekmall/app/payment/biz/repository/dao"
	"github.com/crazyfrankie/seekmall/app/payment/biz/service/wechat"
	"github.com/crazyfrankie/seekmall/rpc_gen/payment"
)

type PaymentServer struct {
	repo   *repository.PaymentRepo
	paysvc wechat.PaymentService

	payment.UnimplementedPaymentServiceServer
}

func NewPaymentServer(repo *repository.PaymentRepo, paysvc wechat.PaymentService) *PaymentServer {
	return &PaymentServer{repo: repo, paysvc: paysvc}
}

func (s *PaymentServer) PrePay(ctx context.Context, req *payment.PrePayRequest) (*payment.PrePayResponse, error) {
	err := s.repo.AddPayment(ctx, &dao.Payment{
		Amt:         req.GetTotal(),
		Currency:    req.GetCurrency(),
		Description: req.GetDescription(),
		BizTradeNo:  "",
		Status:      int8(req.GetStatus()),
	})
	if err != nil {
		return nil, err
	}

	codeUrl, err := s.paysvc.PrePay(ctx, wechat.PayInfo{
		AppID:       req.GetAppId(),
		MchID:       req.GetMchId(),
		Description: req.GetDescription(),
		OutTradeNo:  "",
		NotifyUrl:   "",
		Amount: &wechat.Amount{
			Total:    req.GetTotal(),
			Currency: req.GetCurrency(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &payment.PrePayResponse{CodeUrl: codeUrl}, nil
}

func (s *PaymentServer) HandleCallBack(ctx context.Context, req *payment.HandleCallBackRequest) (*payment.HandleCallBackResponse, error) {
	// 更新数据库状态
	err := s.repo.UpdatePaymentStatus(ctx, &dao.Payment{
		BizTradeNo: req.GetBizTradeNo(),
		TxnID: sql.NullString{
			String: req.GetTransactionId(),
			Valid:  req.GetTransactionId() != "",
		},
		Status: int8(req.GetStatus()),
	})
	if err != nil {
		return nil, err
	}

	return &payment.HandleCallBackResponse{}, nil
}
