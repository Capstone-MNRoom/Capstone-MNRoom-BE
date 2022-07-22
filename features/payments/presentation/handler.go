package presentation

import "be9/mnroom/features/payments"

type PaymentHandler struct {
	paymentBusiness payments.Business
}

func NewPaymentHandler(business payments.Business) *PaymentHandler {
	return &PaymentHandler{
		paymentBusiness: business,
	}
}
