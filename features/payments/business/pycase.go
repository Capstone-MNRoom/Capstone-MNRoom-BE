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

func (py *paymentCase) GetAllData(idToken int) (data []payments.Core, err error) {
	data, err = py.paymentData.GetAllData(idToken)
	return data, err
}

func (py *paymentCase) UpdateData(idOrder string, status string) (row int, err error) {
	row, err = py.paymentData.UpdateData(idOrder, status)
	return row, err
}
