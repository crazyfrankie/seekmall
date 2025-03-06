package domain

type PaymentStatus int8

const (
	PaymentStatusUnknown PaymentStatus = iota
	PaymentStatusInit
	PaymentStatusSuccess
	PaymentStatusFailed
	PaymentStatusRefund
)

func (p PaymentStatus) AsInt8() int8 {
	return int8(p)
}
