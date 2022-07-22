package midtrans

type ChargeResponse struct {
	TransactionID     string     `json:"transaction_id"`
	OrderID           string     `json:"order_id"`
	GrossAmount       string     `json:"gross_amount"`
	PaymentType       string     `json:"payment_type"`
	TransactionStatus string     `json:"transaction_status"`
	FraudStatus       string     `json:"fraud_status"`
	StatusCode        string     `json:"status_code"`
	Bank              string     `json:"bank"`
	StatusMessage     string     `json:"status_message"`
	Currency          string     `json:"currency"`
	VaNumbers         []VANumber `json:"va_numbers"`
}

type VANumber struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number"`
}
