package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/seekmall/app/api/pkg/mws"
	"github.com/crazyfrankie/seekmall/app/api/pkg/response"
	"github.com/crazyfrankie/seekmall/rpc_gen/sm"
	"github.com/crazyfrankie/seekmall/rpc_gen/user"
)

type SendCodeReq struct {
	Phone string `json:"phone"`
}

type VerifyCodeReq struct {
	Biz   string `json:"biz"`
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type UserInfoResp struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UserHandler struct {
	userClient user.UserServiceClient
	sms        sm.SmsServiceClient
}

func NewUserHandler(userClient user.UserServiceClient, sms sm.SmsServiceClient) *UserHandler {
	return &UserHandler{userClient: userClient, sms: sms}
}

func (h *UserHandler) RegisterRoute(r *gin.Engine) {
	userGroup := r.Group("api/user")
	{
		userGroup.POST("send-code", h.SendCode())
		userGroup.POST("verify-code", h.VerifyCode())
		userGroup.GET("", h.UserInfo())
	}
}

func (h *UserHandler) SendCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SendCodeReq
		if err := c.Bind(&req); err != nil {
			return
		}

		resp, err := h.userClient.SendCode(c.Request.Context(), &user.SendCodeRequest{
			Phone: req.Phone,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}

func (h *UserHandler) VerifyCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req VerifyCodeReq
		if err := c.Bind(&req); err != nil {
			return
		}

		resp, err := h.userClient.VerifyCode(c.Request.Context(), &user.VerifyCodeRequest{
			Biz:   req.Biz,
			Phone: req.Phone,
			Code:  req.Code,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		c.Header("x-jwt-token", resp.Token)

		response.Success(c, nil)
	}
}

func (h *UserHandler) UserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims")
		claim, _ := claims.(*mws.Claims)

		resp, err := h.userClient.GetUserInfo(c.Request.Context(), &user.GetUserInfoRequest{
			Uid: int32(claim.UId),
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, UserInfoResp{
			ID:    int(resp.User.Id),
			Name:  resp.User.Name,
			Phone: resp.User.Phone,
		})
	}
}
