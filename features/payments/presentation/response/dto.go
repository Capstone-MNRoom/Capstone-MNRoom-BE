package response

import (
	"be9/mnroom/features/payments"
	_rents "be9/mnroom/features/rents/presentation/response"
	_users "be9/mnroom/features/users/presentation/response"
	"time"
)

type Payments struct {
	ID                int       `json:"id"`
	TransactionID     string    `json:"transaction_id"`
	PaymentType       string    `json:"payment_type"`
	OrderID           int       `json:"order_id"`
	BankTransfer      string    `json:"bank_transfer"`
	GrossAmount       int       `json:"gross_amount"`
	VANumber          int       `json:"va_number"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
	User              _users.User
	Rents             _rents.Rents
}

func FromCoreList(data []payments.Core) []Payments {
	result := []Payments{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data payments.Core) Payments {
	return Payments{
		ID:                data.ID,
		TransactionID:     data.TransactionID,
		PaymentType:       data.PaymentType,
		OrderID:           data.OrderID,
		BankTransfer:      data.BankTransfer,
		GrossAmount:       data.GrossAmount,
		VANumber:          data.VANumber,
		TransactionStatus: data.TransactionStatus,
		CreatedAt:         data.CreatedAt,
		User: _users.User{
			ID:       data.User.ID,
			Username: data.User.Username,
			Email:    data.User.Email,
		},
		Rents: _rents.Rents{
			ID:               data.Rents.ID,
			DateStart:        data.Rents.DateStart,
			DateEnd:          data.Rents.DateEnd,
			TotalRentalPrice: data.Rents.TotalRentalPrice,
		},
	}
}
