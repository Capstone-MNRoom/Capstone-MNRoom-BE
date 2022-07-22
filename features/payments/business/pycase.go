package business

import "be9/mnroom/features/payments"

type paymentCase struct {
	paymentData payments.Data
}

func NewPaymentBusiness(payment payments.Data) payments.Business {
	return &paymentCase{
		paymentData: payment,
	}
}
