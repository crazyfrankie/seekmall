package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"

	"github.com/crazyfrankie/seekmall/app/payment/biz/domain"
	"github.com/crazyfrankie/seekmall/app/payment/biz/repository"
	"github.com/crazyfrankie/seekmall/app/payment/biz/repository/dao"
	"github.com/crazyfrankie/seekmall/app/payment/biz/service/wechat"
	"github.com/crazyfrankie/seekmall/rpc_gen/payment"
)

type PaymentServer struct {
	repo   *repository.PaymentRepo
	paysvc *wechat.NativePayService

	nativeCBTypeToStatus map[string]domain.PaymentStatus

	payment.UnimplementedPaymentServiceServer
}

func NewPaymentServer(repo *repository.PaymentRepo, paysvc *wechat.NativePayService) *PaymentServer {
	status := map[string]domain.PaymentStatus{
		"SUCCESS":    domain.PaymentStatusSuccess, // 支付成功
		"REFUND":     domain.PaymentStatusRefund,  // 转入退款
		"NOTPAY":     domain.PaymentStatusInit,    // 未支付
		"CLOSED":     domain.PaymentStatusFailed,  // 已关闭
		"REVOKED":    domain.PaymentStatusFailed,  // 已撤销(付款码支付)
		"PAYERROR":   domain.PaymentStatusFailed,  // 支付失败(其他原因, 如银行返回失败)
		"USERPAYING": domain.PaymentStatusRefund,  // 用户支付中
	}

	return &PaymentServer{repo: repo, paysvc: paysvc, nativeCBTypeToStatus: status}
}

func (s *PaymentServer) PrePay(ctx context.Context, req *payment.PrePayRequest) (*payment.PrePayResponse, error) {
	err := s.repo.AddPayment(ctx, &dao.Payment{
		Amt:         req.GetTotal(),
		Currency:    req.GetCurrency(),
		Description: req.GetDescription(),
		BizTradeNo:  req.GetBizTradeNo(),
		Status:      domain.PaymentStatusInit.AsInt8(),
	})
	if err != nil {
		return nil, err
	}

	codeUrl, err := s.paysvc.PrePay(ctx, wechat.PayInfo{
		Description: req.GetDescription(),
		OutTradeNo:  req.GetBizTradeNo(),
		NotifyUrl:   req.GetNotifyUrl(),
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
	txn := new(payments.Transaction)
	err := json.Unmarshal(req.GetTransaction(), txn)
	if err != nil {
		return nil, err
	}

	// 更新数据库状态
	err = s.updateByTxn(ctx, txn)
	if err != nil {
		return nil, err
	}

	return &payment.HandleCallBackResponse{}, nil
}

func (s *PaymentServer) FindExpirePayment(ctx context.Context, req *payment.FindExpirePaymentRequest) (*payment.FindExpirePaymentResponse, error) {
	pms, err := s.repo.FindExpirePayment(ctx, int(req.GetLimit()), int(req.GetOffset()), req.GetTime())
	if err != nil {
		return nil, err
	}

	payments := make([]*payment.Payment, len(pms))
	for _, p := range pms {
		payments = append(payments, &payment.Payment{BizTradeNo: p.BizTradeNo})
	}

	return &payment.FindExpirePaymentResponse{Payments: payments}, nil
}

func (s *PaymentServer) SyncWechatInfo(ctx context.Context, req *payment.SyncWechatInfoRequest) (*payment.SyncWechatInfoResponse, error) {
	trades := req.GetBizTradeNo()
	for _, t := range trades {
		txn, err := s.paysvc.QueryOrderByOutTradeNo(ctx, t)
		if err != nil {
			return nil, err
		}

		err = s.updateByTxn(ctx, txn)
		if err != nil {
			return nil, err
		}
	}

	return &payment.SyncWechatInfoResponse{}, nil
}

func (s *PaymentServer) updateByTxn(ctx context.Context, txn *payments.Transaction) error {
	status, ok := s.nativeCBTypeToStatus[*txn.TradeState]
	if !ok {
		return errors.New("未知的回调")
	}

	err := s.repo.UpdatePaymentStatus(ctx, &dao.Payment{
		BizTradeNo: *txn.OutTradeNo,
		TxnID: sql.NullString{
			String: *txn.TransactionId,
			Valid:  *txn.TransactionId != "",
		},
		Status: status.AsInt8(),
	})
	if err != nil {
		return err
	}

	return nil
}
