package service

import (
	"context"

	"github.com/crazyfrankie/gem/gerrors"
	"google.golang.org/grpc"

	"github.com/crazyfrankie/seekmall/app/cart/biz/repository"
	"github.com/crazyfrankie/seekmall/app/cart/biz/repository/dao"
	"github.com/crazyfrankie/seekmall/rpc_gen/cart"
	"github.com/crazyfrankie/seekmall/rpc_gen/product"
)

type CartServer struct {
	repo          *repository.CartRepo
	productClient product.ProductServiceClient
	cart.UnimplementedCartServiceServer
}

func NewCartServer(repo *repository.CartRepo, productClient product.ProductServiceClient) *CartServer {
	return &CartServer{repo: repo, productClient: productClient}
}

func (s *CartServer) RegisterServer(server *grpc.Server) {
	cart.RegisterCartServiceServer(server, s)
}

func (s *CartServer) AddItem(ctx context.Context, req *cart.AddCartRequest) (*cart.AddCartResponse, error) {
	resp, err := s.productClient.GetProduct(ctx, &product.GetProductRequest{
		Id: req.GetPid(),
	})
	if err != nil {
		return nil, err
	}
	if resp.Product.Id == 0 {
		return nil, gerrors.NewBizError(40000, "product not found")
	}

	item := &dao.Item{
		UserId:    req.GetUid(),
		ProductId: req.GetPid(),
		Quantity:  req.GetQuantity(),
	}

	err = s.repo.CreateItem(ctx, item)
	if err != nil {
		return nil, err
	}

	return &cart.AddCartResponse{}, nil
}

func (s *CartServer) CartList(ctx context.Context, req *cart.CartListRequest) (*cart.CartListResponse, error) {
	items, err := s.repo.FindByUserId(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	res := make([]*cart.Item, 0, len(items))
	for _, item := range items {
		res = append(res, &cart.Item{
			Id:        item.Id,
			ProductId: item.ProductId,
			UserId:    item.UserId,
			Quantity:  item.Quantity,
		})
	}

	return &cart.CartListResponse{Items: res}, nil
}

func (s *CartServer) EmptyCart(ctx context.Context, req *cart.EmptyCartRequest) (*cart.EmptyCartResponse, error) {
	err := s.repo.DeleteCartsById(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &cart.EmptyCartResponse{}, nil
}
