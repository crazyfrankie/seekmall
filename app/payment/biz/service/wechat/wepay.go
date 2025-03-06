package wechat

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
)

type PaymentService interface {
	PrePay(ctx context.Context, info PayInfo) (string, error)
}

type NativePayService struct {
	svc *native.NativeApiService
}

type PayInfo struct {
	AppID       string
	MchID       string
	Description string
	OutTradeNo  string
	NotifyUrl   string
	Amount      *Amount
}

type Amount struct {
	Total    int64
	Currency string
}

func NewNativePayService(svc *native.NativeApiService) PaymentService {
	return &NativePayService{svc: svc}
}

func (p *NativePayService) PrePay(ctx context.Context, info PayInfo) (string, error) {
	resp, _, err := p.svc.Prepay(ctx, native.PrepayRequest{
		Appid:       toPtr[string](info.AppID),
		Mchid:       toPtr[string](info.MchID),
		Description: toPtr[string](info.Description),
		OutTradeNo:  toPtr[string](info.OutTradeNo),
		NotifyUrl:   toPtr[string](info.NotifyUrl),
		Amount: &native.Amount{
			Total:    toPtr[int64](info.Amount.Total),
			Currency: toPtr[string](info.Amount.Currency),
		}})
	if err != nil {
		return "", err
	}

	return *resp.CodeUrl, nil
}

func toPtr[T any](s T) *T {
	return &s
}
