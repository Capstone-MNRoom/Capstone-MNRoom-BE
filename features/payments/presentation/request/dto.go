package request

import (
	"be9/mnroom/features/payments"
	"be9/mnroom/features/rents"
)

type Payments struct {
	TransactionID     string `json:"transaction_id" form:"transaction_id"`
	PaymentType       string `json:"payment_type" form:"payment_type"`
	OrderID           int    `json:"order_id" form:"order_id"`
	RentsID           int    `json:"rents_id" form:"rents_id"`
	BankTransfer      string `json:"bank_transfer" form:"bank_transfer"`
	GrossAmount       int    `json:"gross_amount" form:"gross_amount"`
	VANumber          int    `json:"va_number" form:"va_number"`
	TransactionStatus string `json:"transaction_status" form:"transaction_status"`
}

func ToCore(req Payments) payments.Core {
	return payments.Core{
		TransactionID: req.TransactionID,
		PaymentType:   req.PaymentType,
		OrderID:       req.OrderID,
		Rents: rents.Core{
			ID: req.RentsID,
		},
		BankTransfer:      req.BankTransfer,
		GrossAmount:       req.GrossAmount,
		VANumber:          req.VANumber,
		TransactionStatus: req.TransactionStatus,
	}
}
