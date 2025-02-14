package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	Id        int32 `gorm:"primaryKey,autoIncrement"`
	ProductId int32
	UserId    int32 `gorm:"index:uid"`
	Quantity  int32
	Ctime     int64
	Utime     int64
}

type CartDao struct {
	db *gorm.DB
}

func NewCartDao(db *gorm.DB) *CartDao {
	return &CartDao{db: db}
}

func (d *CartDao) CreateItem(ctx context.Context, i *Item) error {
	now := time.Now().Unix()
	i.Ctime = now
	i.Utime = now

	return d.db.WithContext(ctx).Model(&Item{}).Create(i).Error
}

func (d *CartDao) FindByUserId(ctx context.Context, uid int32) ([]*Item, error) {
	var items []*Item
	err := d.db.WithContext(ctx).Model(&Item{}).Where("user_id = ?", uid).Find(&items).Error
	if err != nil {
		return []*Item{}, err
	}

	return items, nil
}

func (d *CartDao) DeleteCartsById(ctx context.Context, uid int32) error {
	return d.db.WithContext(ctx).Model(&Item{}).Delete("user_id = ?", uid).Error
}
