package repository

import (
	"context"
	"github.com/crazyfrankie/seekmall/app/payment/biz/repository/dao"
)

type PaymentRepo struct {
	dao *dao.PaymentDao
}

func NewPaymentRepo(d *dao.PaymentDao) *PaymentRepo {
	return &PaymentRepo{dao: d}
}

func (r *PaymentRepo) AddPayment(ctx context.Context, p *dao.Payment) error {
	return r.dao.CreatePay(ctx, p)
}

func (r *PaymentRepo) UpdatePaymentStatus(ctx context.Context, p *dao.Payment) error {
	return r.dao.UpdatePaymentStatus(ctx, p)
}
