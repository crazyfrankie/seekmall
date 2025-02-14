package handler

import (
	"github.com/crazyfrankie/seekmall/app/api/pkg/mws"
	"github.com/crazyfrankie/seekmall/app/api/pkg/response"
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/seekmall/rpc_gen/cart"
)

type AddItemReq struct {
	ProductId int32 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type CartHandler struct {
	client cart.CartServiceClient
}

func NewCartHandler(client cart.CartServiceClient) *CartHandler {
	return &CartHandler{client: client}
}

func (h *CartHandler) RegisterRoute(r *gin.Engine) {
	cartGroup := r.Group("api/cart")
	{
		cartGroup.POST("/add", h.AddItem())
		cartGroup.GET("", h.CartList())
		cartGroup.DELETE("", h.EmptyCart())
	}
}

func (h *CartHandler) AddItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddItemReq
		if err := c.Bind(&req); err != nil {
			return
		}

		claims := c.MustGet("claims")
		claim, _ := claims.(*mws.Claims)

		resp, err := h.client.AddCart(c.Request.Context(), &cart.AddCartRequest{
			Pid:      req.ProductId,
			Quantity: req.Quantity,
			Uid:      claim.UId,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}

func (h *CartHandler) CartList() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims")
		claim, _ := claims.(*mws.Claims)

		resp, err := h.client.CartList(c.Request.Context(), &cart.CartListRequest{
			UserId: claim.UId,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}

func (h *CartHandler) EmptyCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims")
		claim, _ := claims.(*mws.Claims)

		resp, err := h.client.EmptyCart(c.Request.Context(), &cart.EmptyCartRequest{
			UserId: claim.UId,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}
