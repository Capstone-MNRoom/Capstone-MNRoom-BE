package response

import (
	"be9/mnroom/features/payments"
	"time"
)

type Payments struct {
	ID                int       `json:"id"`
	TransactionID     string    `json:"transaction_id"`
	PaymentType       string    `json:"payment_type"`
	OrderID           string    `json:"order_id"`
	BankTransfer      string    `json:"bank_transfer"`
	GrossAmount       int       `json:"gross_amount"`
	VANumber          string    `json:"va_number"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
}

type PaymentsCore struct {
	ID                int       `json:"id"`
	TransactionID     string    `json:"transaction_id"`
	PaymentType       string    `json:"payment_type"`
	OrderID           string    `json:"order_id"`
	BankTransfer      string    `json:"bank_transfer"`
	GrossAmount       int       `json:"gross_amount"`
	VANumber          string    `json:"va_number"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
	User              User
	Rents             Rents
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Rents struct {
	ID               int    `json:"id"`
	DateStart        string `json:"date_start"`
	DateEnd          string `json:"date_end"`
	Bank             string `json:"bank"`
	TotalRentalPrice int    `json:"total_rental_price"`
	Status           string `json:"status"`
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
	}
}

func FromCoreListPayments(data []payments.Core) []PaymentsCore {
	result := []PaymentsCore{}
	for key := range data {
		result = append(result, FromCorePayments(data[key]))
	}
	return result
}

func FromCorePayments(data payments.Core) PaymentsCore {
	return PaymentsCore{
		ID:                data.ID,
		TransactionID:     data.TransactionID,
		PaymentType:       data.PaymentType,
		OrderID:           data.OrderID,
		BankTransfer:      data.BankTransfer,
		GrossAmount:       data.GrossAmount,
		VANumber:          data.VANumber,
		TransactionStatus: data.TransactionStatus,
		CreatedAt:         data.CreatedAt,
		User: User{
			ID:       data.User.ID,
			Username: data.User.Username,
			Email:    data.User.Email,
		},
		Rents: Rents{
			ID:               data.Rents.ID,
			DateStart:        data.Rents.DateStart,
			DateEnd:          data.Rents.DateEnd,
			Bank:             data.Rents.Bank,
			TotalRentalPrice: data.Rents.TotalRentalPrice,
			Status:           data.Rents.Status,
		},
	}
}
