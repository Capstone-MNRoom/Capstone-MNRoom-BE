package business

import (
	"be9/mnroom/features/payments"
	"context"
)

type paymentCase struct {
	paymentData payments.Data
}

func NewPaymentBusiness(payment payments.Data) payments.Business {
	return &paymentCase{
		paymentData: payment,
	}
}

func (py *paymentCase) GetAllData(ctx context.Context, idToken int) (data []payments.Core, err error) {
	data, err = py.paymentData.GetAllData(ctx, idToken)
	return data, err
}

func (py *paymentCase) UpdateData(ctx context.Context, idOrder string, status string) (row int, err error) {
	row, err = py.paymentData.UpdateData(ctx, idOrder, status)
	return row, err
}
