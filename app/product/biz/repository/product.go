package repository

import (
	"context"

	"github.com/crazyfrankie/seekmall/app/product/biz/repository/dao"
)

type ProductRepo struct {
	seller    *dao.SellerDao
	purchaser *dao.PurchaserDao
}

func NewProductRepo(seller *dao.SellerDao, purchaser *dao.PurchaserDao) *ProductRepo {
	return &ProductRepo{seller: seller, purchaser: purchaser}
}

func (r *ProductRepo) CreateDraft(ctx context.Context, p *dao.ProductDraft) error {
	return r.seller.CreateDraft(ctx, p)
}

func (r *ProductRepo) QueryDraftExist(ctx context.Context, name string) (bool, error) {
	return r.seller.QueryDraftExist(ctx, name)
}

func (r *ProductRepo) QueryDraftById(ctx context.Context, pid int) (bool, error) {
	return r.seller.QueryDraftById(ctx, pid)
}

func (r *ProductRepo) SyncToLive(ctx context.Context, pid int) error {
	return r.seller.CreateLive(ctx, pid)
}

func (r *ProductRepo) GetProductById(ctx context.Context, pid int) (dao.ProductLive, error) {
	return r.purchaser.GetProductById(ctx, pid)
}

func (r *ProductRepo) GetProductsByCategoryName(ctx context.Context, limit, offset int, name string) ([]*dao.Category, error) {
	return r.purchaser.GetProductsByCategoryName(ctx, limit, offset, name)
}

func (r *ProductRepo) SearchProducts(ctx context.Context, query string) ([]*dao.ProductLive, error) {
	return r.purchaser.SearchProducts(ctx, query)
}
