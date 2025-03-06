package wechat

import (
	"context"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
)

type NativePayService struct {
	svc   *native.NativeApiService
	appId string
	mchId string
}

type PayInfo struct {
	Description string
	OutTradeNo  string
	NotifyUrl   string
	Amount      *Amount
}

type Amount struct {
	Total    int64
	Currency string
}

func NewNativePayService(svc *native.NativeApiService) *NativePayService {
	return &NativePayService{svc: svc, appId: "", mchId: ""}
}

func (p *NativePayService) PrePay(ctx context.Context, info PayInfo) (string, error) {
	resp, _, err := p.svc.Prepay(ctx, native.PrepayRequest{
		Appid:       toPtr[string](p.appId),
		Mchid:       toPtr[string](p.mchId),
		Description: toPtr[string](info.Description),
		OutTradeNo:  toPtr[string](info.OutTradeNo),
		NotifyUrl:   toPtr[string](info.NotifyUrl),
		Amount: &native.Amount{
			Total:    toPtr[int64](info.Amount.Total),
			Currency: toPtr[string](info.Amount.Currency),
		},
		TimeExpire: toPtr[time.Time](time.Now().Add(time.Minute * 30)),
	})
	if err != nil {
		return "", err
	}

	return *resp.CodeUrl, nil
}

func (p *NativePayService) QueryOrderByOutTradeNo(ctx context.Context, outTradeNo string) (*payments.Transaction, error) {
	txn, _, err := p.svc.QueryOrderByOutTradeNo(ctx, native.QueryOrderByOutTradeNoRequest{
		OutTradeNo: toPtr[string](outTradeNo),
		Mchid:      toPtr[string](p.mchId),
	})
	if err != nil {
		return nil, err
	}

	return txn, nil
}

func toPtr[T any](s T) *T {
	return &s
}
