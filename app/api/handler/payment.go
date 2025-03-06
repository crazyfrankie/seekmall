package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"

	"github.com/crazyfrankie/seekmall/app/api/domain"
	"github.com/crazyfrankie/seekmall/app/api/pkg/response"
	"github.com/crazyfrankie/seekmall/rpc_gen/payment"
)

type PaymentHandler struct {
	handler *notify.Handler
	client  payment.PaymentServiceClient

	nativeCBTypeToStatus map[string]domain.PaymentStatus
}

func NewPaymentHandler(handler *notify.Handler, client payment.PaymentServiceClient) *PaymentHandler {
	status := map[string]domain.PaymentStatus{
		"SUCCESS":    domain.PaymentStatusSuccess, // 支付成功
		"REFUND":     domain.PaymentStatusRefund,  // 转入退款
		"NOTPAY":     domain.PaymentStatusInit,    // 未支付
		"CLOSED":     domain.PaymentStatusFailed,  // 已关闭
		"REVOKED":    domain.PaymentStatusFailed,  // 已撤销(付款码支付)
		"PAYERROR":   domain.PaymentStatusFailed,  // 支付失败(其他原因, 如银行返回失败)
		"USERPAYING": domain.PaymentStatusRefund,  // 用户支付中
	}

	return &PaymentHandler{
		handler:              handler,
		client:               client,
		nativeCBTypeToStatus: status,
	}
}

func (h *PaymentHandler) RegisterRoute(r *gin.Engine) {
	payGroup := r.Group("api/payment")
	{
		payGroup.Any("/wechat/callback", h.NativeHandler())
	}
}

func (h *PaymentHandler) NativeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var txn payments.Transaction
		_, err := h.handler.ParseNotifyRequest(c.Request.Context(), c.Request, &txn)
		if err != nil {
			response.Error(c, err)
			return
		}

		status, ok := h.nativeCBTypeToStatus[*txn.TradeState]
		if !ok {
			response.Error(c, errors.New("未知的回调"))
			return
		}

		resp, err := h.client.HandleCallBack(c.Request.Context(), &payment.HandleCallBackRequest{
			BizTradeNo:    *txn.OutTradeNo,
			TransactionId: *txn.TransactionId,
			Status:        int32(status.AsInt8()),
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}
