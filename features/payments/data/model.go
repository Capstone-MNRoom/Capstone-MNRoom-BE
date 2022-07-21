package data

import (
	"be9/mnroom/features/payments"
	"be9/mnroom/features/rents"
	_rents "be9/mnroom/features/rents/data"
	"be9/mnroom/features/users"
	_users "be9/mnroom/features/users/data"

	"gorm.io/gorm"
)

type Payments struct {
	gorm.Model
	UserID            int    `json:"user_id" form:"user_id"`
	TransactionID     string `json:"transaction_id" form:"transaction_id"`
	PaymentType       string `json:"payment_type" form:"payment_type"`
	OrderID           int    `json:"order_id" form:"order_id"`
	RentsID           int    `json:"rents_id" form:"rents_id"`
	BankTransfer      string `json:"bank_transfer" form:"bank_transfer"`
	GrossAmount       int    `json:"gross_amount" form:"gross_amount"`
	VANumber          int    `json:"va_number" form:"va_number"`
	TransactionStatus string `json:"transaction_status" form:"transaction_status"`
	User              _users.User
	Rents             _rents.Rents
}

func toCoreList(data []Payments) []payments.Core {
	result := []payments.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *Payments) toCore() payments.Core {
	return payments.Core{
		ID:                int(data.ID),
		TransactionID:     data.TransactionID,
		PaymentType:       data.PaymentType,
		OrderID:           data.OrderID,
		BankTransfer:      data.BankTransfer,
		GrossAmount:       data.GrossAmount,
		VANumber:          data.VANumber,
		TransactionStatus: data.TransactionStatus,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
		User: users.Core{
			ID:       int(data.User.ID),
			Username: data.User.Username,
			Email:    data.User.Email,
		},
		Rents: rents.Core{
			ID:               int(data.Rents.ID),
			DateStart:        data.Rents.DateStart,
			DateEnd:          data.Rents.DateEnd,
			TotalRentalPrice: data.Rents.TotalRentalPrice,
		},
	}
}

func fromCore(core payments.Core) Payments {
	return Payments{
		TransactionID:     core.TransactionID,
		PaymentType:       core.PaymentType,
		OrderID:           core.OrderID,
		BankTransfer:      core.BankTransfer,
		GrossAmount:       core.GrossAmount,
		VANumber:          core.VANumber,
		TransactionStatus: core.TransactionStatus,
		UserID:            core.User.ID,
		RentsID:           core.Rents.ID,
	}
}
