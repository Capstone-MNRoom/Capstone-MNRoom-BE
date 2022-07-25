package data

import (
	"be9/mnroom/features/payments"
	"be9/mnroom/features/rents"
	"be9/mnroom/features/users"

	"gorm.io/gorm"
)

type Payments struct {
	gorm.Model
	TransactionID     string `json:"transaction_id" form:"transaction_id"`
	PaymentType       string `json:"payment_type" form:"payment_type"`
	OrderID           string `json:"order_id" form:"order_id"`
	BankTransfer      string `json:"bank_transfer" form:"bank_transfer"`
	GrossAmount       int    `json:"gross_amount" form:"gross_amount"`
	VANumber          string `json:"va_number" form:"va_number"`
	TransactionStatus string `json:"transaction_status" form:"transaction_status"`
	UserID            int    `json:"user_id" form:"user_id"`
	RoomsID           int    `json:"rooms_id" form:"rooms_id"`
	RentsID           int    `json:"rents_id" form:"rents_id"`
	User              User
	Rents             Rents
}

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Email    string `gorm:"unique" json:"email" form:"email"`
}

type Rents struct {
	gorm.Model
	DateStart        string `json:"date_start" form:"date_start"`
	DateEnd          string `json:"date_end" form:"date_end"`
	Bank             string `json:"bank" form:"bank"`
	TotalRentalPrice int    `json:"total_rental_price" form:"total_rental_price"`
	Status           string `json:"status" form:"status"`
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
			Bank:             data.Rents.Bank,
			TotalRentalPrice: data.Rents.TotalRentalPrice,
			Status:           data.Rents.Status,
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
