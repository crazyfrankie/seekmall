package dao

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/crazyfrankie/seekmall/app/product/pkg/constants"
)

type SellerDao struct {
	db *gorm.DB
}

func NewSellerDao(db *gorm.DB) *SellerDao {
	return &SellerDao{db: db}
}

func (d *SellerDao) CreateDraft(ctx context.Context, draft *ProductDraft) error {
	now := time.Now().Unix()
	draft.Ctime = now
	draft.Utime = now

	return d.db.WithContext(ctx).Model(&ProductDraft{}).Create(draft).Error
}

func (d *SellerDao) CreateLive(ctx context.Context, id int) error {
	err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var draft ProductDraft
		err := tx.WithContext(ctx).Model(&ProductDraft{}).Where("id = ?", id).Find(&draft).Error
		if err != nil {
			return err
		}
		if draft.Id == 0 {
			return constants.ProductDraftNotFound
		}

		now := time.Now().Unix()
		live := &ProductLive{
			Name:        draft.Name,
			Description: draft.Description,
			Picture:     draft.Picture,
			Price:       draft.Price,
			Uid:         draft.Uid,
			Categories:  draft.Categories,
			Status:      "live",
			Ctime:       now,
			Utime:       now,
		}

		return tx.WithContext(ctx).Model(&ProductLive{}).Create(live).Error
	})

	return err
}

func (d *SellerDao) QueryDraftExist(ctx context.Context, name string) (bool, error) {
	var draft ProductDraft
	err := d.db.WithContext(ctx).Model(&ProductDraft{}).Where("name = ?", name).Find(&draft).Error
	if err != nil {
		return false, err
	}
	if draft.Id == 0 {
		return false, nil
	}

	return true, nil
}

func (d *SellerDao) QueryDraftById(ctx context.Context, pid int) (bool, error) {
	var draft ProductDraft
	err := d.db.WithContext(ctx).Model(&ProductDraft{}).Where("id = ?", pid).Find(&draft).Error
	if err != nil {
		return false, err
	}
	if draft.Id == 0 {
		return false, nil
	}

	return true, nil
}

type PurchaserDao struct {
	db *gorm.DB
}

func NewPurchaserDao(db *gorm.DB) *PurchaserDao {
	return &PurchaserDao{db: db}
}

func (d *PurchaserDao) GetProductById(ctx context.Context, pid int) (ProductLive, error) {
	var p ProductLive
	err := d.db.WithContext(ctx).Model(&ProductLive{}).Where("id = ?", pid).Find(&p).Error
	if err != nil {
		return ProductLive{}, err
	}

	return p, nil
}

func (d *PurchaserDao) GetProductsByCategoryName(ctx context.Context, limit, offset int, name string) ([]*Category, error) {
	var categories []*Category
	err := d.db.WithContext(ctx).Model(&Category{}).
		Where("name = ?", name).Preload("Product").Limit(limit).Offset(offset).
		Find(&categories).Error
	if err != nil {
		return []*Category{}, err
	}

	return categories, nil
}

func (d *PurchaserDao) SearchProducts(ctx context.Context, query string) ([]*ProductLive, error) {
	var products []*ProductLive
	err := d.db.WithContext(ctx).Model(&ProductLive{}).
		Where("name like ? OR description like ?", "%"+query+"%", "%"+query+"%").
		Find(&products).Error
	if err != nil {
		return []*ProductLive{}, err
	}

	return products, nil
}
