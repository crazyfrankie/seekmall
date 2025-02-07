package service

import (
	"context"

	"google.golang.org/grpc"

	"github.com/crazyfrankie/seekmall/app/product/biz/repository"
	"github.com/crazyfrankie/seekmall/app/product/biz/repository/dao"
	"github.com/crazyfrankie/seekmall/app/product/pkg/constants"
	"github.com/crazyfrankie/seekmall/rpc_gen/product"
)

type ProductServer struct {
	repo *repository.ProductRepo
	product.UnimplementedProductServiceServer
}

func NewProductServer(repo *repository.ProductRepo) *ProductServer {
	return &ProductServer{repo: repo}
}

func (s *ProductServer) RegisterServer(server *grpc.Server) {
	product.RegisterProductServiceServer(server, s)
}

func (s *ProductServer) AddProduct(ctx context.Context, req *product.AddProductRequest) (*product.AddProductResponse, error) {
	exists, err := s.repo.QueryDraftExist(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, constants.ProductDraftExists
	}

	cgs := make([]dao.Category, len(req.GetCategories()))
	for _, cg := range req.GetCategories() {
		cgs = append(cgs, dao.Category{
			Name: cg,
		})
	}

	err = s.repo.CreateDraft(ctx, &dao.ProductDraft{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Picture:     req.GetPicture(),
		Price:       req.GetPrice(),
		Uid:         int(req.GetUid()),
		Categories:  cgs,
	})
	if err != nil {
		return nil, err
	}

	return &product.AddProductResponse{}, nil
}

func (s *ProductServer) ReleaseProduct(ctx context.Context, req *product.ReleaseProductRequest) (*product.ReleaseProductResponse, error) {
	pid := int(req.GetId())

	exists, err := s.repo.QueryDraftById(ctx, pid)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, constants.ProductDraftNotFound
	}

	err = s.repo.SyncToLive(ctx, pid)
	if err != nil {
		return nil, err
	}

	return &product.ReleaseProductResponse{}, nil
}

func (s *ProductServer) GetProduct(ctx context.Context, req *product.GetProductRequest) (*product.GetProductResponse, error) {
	p, err := s.repo.GetProductById(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	cgs := make([]string, len(p.Categories))
	for _, cg := range p.Categories {
		cgs = append(cgs, cg.Name)
	}
	return &product.GetProductResponse{
		Product: &product.Product{
			Id:          int32(p.Id),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			Categories:  cgs,
		},
	}, nil
}

func (s *ProductServer) ListProducts(ctx context.Context, req *product.ListProductsRequest) (*product.ListProductsResponse, error) {
	limit := int(req.GetPageSize())
	offset := int((req.GetPage() - 1) * req.GetPageSize())
	cgs, err := s.repo.GetProductsByCategoryName(ctx, limit, offset, req.GetCategoryName())
	if err != nil {
		return nil, err
	}

	var products []*product.Product
	for _, cg := range cgs {
		for _, p := range cg.LiveProducts {
			products = append(products, &product.Product{
				Id:          int32(p.Id),
				Name:        p.Name,
				Description: p.Description,
				Picture:     p.Picture,
				Price:       p.Price,
			})
		}
	}

	return &product.ListProductsResponse{
		Products: products,
	}, nil
}

func (s *ProductServer) SearchProducts(ctx context.Context, req *product.SearchProductsRequest) (*product.SearchProductsResponse, error) {
	resp, err := s.repo.SearchProducts(ctx, req.GetQuery())
	if err != nil {
		return nil, err
	}

	var products []*product.Product
	for _, p := range resp {
		products = append(products, &product.Product{
			Id:          int32(p.Id),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
		})
	}

	return &product.SearchProductsResponse{Results: products}, nil
}
