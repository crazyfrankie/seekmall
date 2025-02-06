package repository

import (
	"context"

	"github.com/crazyfrankie/seekmall/app/product/domain"
	"github.com/crazyfrankie/seekmall/app/product/repository/dao"
)

type ProductRepo struct {
	seller    *dao.SellerDao
	purchaser *dao.PurchaserDao
}

func NewProductRepo(seller *dao.SellerDao, purchaser *dao.PurchaserDao) *ProductRepo {
	return &ProductRepo{seller: seller, purchaser: purchaser}
}

func (r *ProductRepo) CreateDraft(ctx context.Context, p domain.Product) error {
	return r.seller.CreateDraft(ctx, r.domainToDao(p))
}

func (r *ProductRepo) QueryDraftExist(ctx context.Context, name string, uid int) (bool, error) {
	return r.seller.QueryDraftExist(ctx, name, uid)
}

func (r *ProductRepo) QueryDraftById(ctx context.Context, pid int) (bool, error) {
	return r.seller.QueryDraftById(ctx, pid)
}

func (r *ProductRepo) SyncToLive(ctx context.Context, pid int) error {
	return r.seller.CreateLive(ctx, pid)
}

func (r *ProductRepo) GetProductById(ctx context.Context, pid int) (domain.Product, error) {
	p, err := r.purchaser.GetProductById(ctx, pid)
	return r.daoToDomain(p), err
}

func (r *ProductRepo) domainToDao(p domain.Product) *dao.ProductDraft {
	categories := make([]dao.Category, len(p.Categories))
	for _, ca := range p.Categories {
		categories = append(categories, dao.Category{
			Name: ca,
		})
	}
	return &dao.ProductDraft{
		Name:        p.Name,
		Description: p.Description,
		Picture:     p.Picture,
		Price:       p.Price,
		Uid:         p.Uid,
		Categories:  categories,
	}
}

func (r *ProductRepo) daoToDomain(p dao.ProductLive) domain.Product {
	categories := make([]string, len(p.Categories))
	for _, ca := range p.Categories {
		categories = append(categories, ca.Name)
	}
	return domain.Product{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Picture:     p.Picture,
		Price:       p.Price,
		Uid:         p.Uid,
		Categories:  categories,
	}
}
