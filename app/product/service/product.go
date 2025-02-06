package service

import (
	"context"
	"strconv"

	"github.com/crazyfrankie/seekmall/app/pkg/constants"
	"github.com/crazyfrankie/seekmall/app/product/domain"
	"github.com/crazyfrankie/seekmall/app/product/repository"
)

type ProductService struct {
	repo *repository.ProductRepo
}

func NewProductService(repo *repository.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) AddProduct(ctx context.Context, p domain.Product) error {
	exists, err := s.repo.QueryDraftExist(ctx, p.Name, p.Uid)
	if err != nil {
		return err
	}
	if exists {
		return constants.ProductDraftExists
	}

	return s.repo.CreateDraft(ctx, p)
}

func (s *ProductService) ReleaseProduct(ctx context.Context, pid int) error {
	exists, err := s.repo.QueryDraftById(ctx, pid)
	if err != nil {
		return err
	}
	if !exists {
		return constants.ProductDraftNotFound
	}

	return s.repo.SyncToLive(ctx, pid)
}

func (s *ProductService) GetProduct(ctx context.Context, id string) (domain.Product, error) {
	pid, _ := strconv.Atoi(id)

	product, err := s.repo.GetProductById(ctx, pid)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
