package handler

import (
	"github.com/crazyfrankie/seekmall/app/pkg/mws"
	"github.com/crazyfrankie/seekmall/app/pkg/response"
	"github.com/gin-gonic/gin"

	ms "github.com/crazyfrankie/seekmall/app/ms/service"
	"github.com/crazyfrankie/seekmall/app/user/service"
)

type SendCodeReq struct {
	Phone string `json:"phone"`
}

type VerifyCodeReq struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type UserInfoResp struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UserHandler struct {
	svc service.UserSvc
	sms ms.SmsSvc
}

func NewUserHandler(svc service.UserSvc, sms ms.SmsSvc) *UserHandler {
	return &UserHandler{svc: svc, sms: sms}
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

		err := h.sms.SendCode(c.Request.Context(), req.Phone)
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, nil)
	}
}

func (h *UserHandler) VerifyCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req VerifyCodeReq
		if err := c.Bind(&req); err != nil {
			return
		}

		err := h.sms.VerifyCode(c.Request.Context(), req.Phone, req.Code)
		if err != nil {
			response.Error(c, err)
			return
		}

		var token string
		token, err = h.svc.FindOrCreateUser(c.Request.Context(), req.Phone)
		if err != nil {
			response.Error(c, err)
			return
		}

		c.Header("x-jwt-token", token)

		response.Success(c, nil)
	}
}

func (h *UserHandler) UserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims")
		claim, _ := claims.(*mws.Claims)

		u, err := h.svc.FindById(c.Request.Context(), claim.UId)
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, UserInfoResp{
			ID:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		})
	}
}
