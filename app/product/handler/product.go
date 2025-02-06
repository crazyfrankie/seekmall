package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/seekmall/app/pkg/mws"
	"github.com/crazyfrankie/seekmall/app/pkg/response"
	"github.com/crazyfrankie/seekmall/app/product/domain"
	"github.com/crazyfrankie/seekmall/app/product/service"
)

type AddProductReq struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Price       float32  `json:"price"`
	Categories  []string `json:"categories"`
}

type ReleaseReq struct {
	ProductId int `json:"product_id"`
}

type ProductResp struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Price       float32  `json:"price"`
	Uid         int      `json:"uid"`
	Categories  []string `json:"categories"`
}

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(svc *service.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) RegisterRoute(r *gin.Engine) {
	productGroup := r.Group("api/product")
	{
		productGroup.POST("", h.AddProduct())
		productGroup.GET("/:id", h.GetProduct())
		productGroup.POST("/release", h.ReleaseProduct())
	}
}

func (h *ProductHandler) AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AddProductReq
		if err := c.Bind(&req); err != nil {
			return
		}

		claims := c.MustGet("claims")
		claim, _ := claims.(*mws.Claims)
		err := h.svc.AddProduct(c.Request.Context(), domain.Product{
			Name:        req.Name,
			Description: req.Description,
			Picture:     req.Picture,
			Price:       req.Price,
			Uid:         claim.UId,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, nil)
	}
}

func (h *ProductHandler) ReleaseProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ReleaseReq
		if err := c.Bind(&req); err != nil {
			return
		}

		err := h.svc.ReleaseProduct(c.Request.Context(), req.ProductId)
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, nil)
	}
}

func (h *ProductHandler) GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		pid := c.Query("id")

		product, err := h.svc.GetProduct(c.Request.Context(), pid)
		if err != nil {
			response.Error(c, err)
			return
		}

		categories := make([]string, len(product.Categories))
		for _, ca := range product.Categories {
			categories = append(categories, ca)
		}

		response.Success(c, ProductResp{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Picture:     product.Picture,
			Price:       product.Price,
			Uid:         product.Uid,
			Categories:  categories,
		})
	}
}
