package handler

import (
	"strconv"
	
	"github.com/gin-gonic/gin"

	"github.com/crazyfrankie/seekmall/app/api/pkg/mws"
	"github.com/crazyfrankie/seekmall/app/api/pkg/response"
	"github.com/crazyfrankie/seekmall/rpc_gen/product"
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

type ListProductReq struct {
	Page         int    `json:"page"`
	PageSize     int    `json:"page_size"`
	CategoryName string `json:"category_name"`
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
	client product.ProductServiceClient
}

func NewProductHandler(client product.ProductServiceClient) *ProductHandler {
	return &ProductHandler{client: client}
}

func (h *ProductHandler) RegisterRoute(r *gin.Engine) {
	productGroup := r.Group("api/product")
	{
		productGroup.POST("", h.AddProduct())
		productGroup.GET("/:id", h.GetProduct())
		productGroup.POST("/release", h.ReleaseProduct())
		productGroup.POST("/list", h.ListProducts())
		productGroup.GET("/search", h.SearchProducts())
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
		_, err := h.client.AddProduct(c.Request.Context(), &product.AddProductRequest{
			Name:        req.Name,
			Description: req.Description,
			Picture:     req.Picture,
			Price:       req.Price,
			Uid:         int32(claim.UId),
			Categories:  req.Categories,
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

		_, err := h.client.ReleaseProduct(c.Request.Context(), &product.ReleaseProductRequest{
			Id: int32(req.ProductId),
		})
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
		id, _ := strconv.Atoi(pid)
		resp, err := h.client.GetProduct(c.Request.Context(), &product.GetProductRequest{
			Id: int32(id),
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		categories := make([]string, len(resp.Product.Categories))
		for _, ca := range resp.Product.Categories {
			categories = append(categories, ca)
		}

		response.Success(c, ProductResp{
			Id:          int(resp.Product.Id),
			Name:        resp.Product.Name,
			Description: resp.Product.Description,
			Picture:     resp.Product.Picture,
			Price:       resp.Product.Price,
			Uid:         int(resp.Product.Uid),
			Categories:  categories,
		})
	}
}

func (h *ProductHandler) ListProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ListProductReq
		if err := c.Bind(&req); err != nil {
			return
		}

		resp, err := h.client.ListProducts(c.Request.Context(), &product.ListProductsRequest{
			Page:         int32(req.Page),
			PageSize:     int32(req.PageSize),
			CategoryName: req.CategoryName,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}

func (h *ProductHandler) SearchProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("query")

		resp, err := h.client.SearchProducts(c.Request.Context(), &product.SearchProductsRequest{
			Query: query,
		})
		if err != nil {
			response.Error(c, err)
			return
		}

		response.Success(c, resp)
	}
}
