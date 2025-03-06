package domain

type PaymentStatus int8

const (
	PaymentStatusRefund PaymentStatus = iota
	PaymentStatusInit
	PaymentStatusSuccess
	PaymentStatusFailed
)

func (p PaymentStatus) AsInt8() int8 {
	return int8(p)
}
