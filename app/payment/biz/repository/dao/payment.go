package dao

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	Id          int64 `gorm:"primaryKey,autoIncrement"`
	Amt         int64
	Currency    string
	Description string
	// 业务方传过来的
	BizTradeNo string `gorm:"column:biz_trade_no;type:varchar(255);unique"`
	// 第三方支付平台的事务 ID
	TxnID  sql.NullString `gorm:"column:biz_trade_no;type:varchar(255);unique"`
	Status int8
	Ctime  int64
	Utime  int64
}

type PaymentDao struct {
	db *gorm.DB
}

func NewPaymentDao(db *gorm.DB) *PaymentDao {
	return &PaymentDao{db: db}
}

func (d *PaymentDao) CreatePay(ctx context.Context, p *Payment) error {
	now := time.Now().UnixMilli()
	p.Ctime = now
	p.Utime = now

	return d.db.WithContext(ctx).Create(p).Error
}

func (d *PaymentDao) UpdatePaymentStatus(ctx context.Context, p *Payment) error {
	return d.db.WithContext(ctx).Model(&Payment{}).
		Where("biz_trade_no = ?", p.BizTradeNo).
		Updates(map[string]any{
			"txn_id": p.TxnID,
			"status": p.Status,
			"utime":  time.Now().UnixMilli(),
		}).Error
}
