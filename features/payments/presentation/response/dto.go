package response

import (
	"be9/mnroom/features/payments"
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
