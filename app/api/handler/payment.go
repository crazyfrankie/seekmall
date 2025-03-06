package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"

	"github.com/crazyfrankie/seekmall/app/api/pkg/response"
	"github.com/crazyfrankie/seekmall/rpc_gen/payment"
)

type PaymentHandler struct {
	handler *notify.Handler
	client  payment.PaymentServiceClient
}

func NewPaymentHandler(handler *notify.Handler, client payment.PaymentServiceClient) *PaymentHandler {
	return &PaymentHandler{handler: handler, client: client}
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
			// 有人伪造请求
			response.Error(c, err)
			return
		}

		data, err := json.Marshal(txn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		resp, err := h.client.HandleCallBack(c.Request.Context(), &payment.HandleCallBackRequest{
			Transaction: data,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		response.Success(c, resp)
	}
}
