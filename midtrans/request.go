package midtrans

import "github.com/midtrans/midtrans-go"

type ChargeReq struct {
	PaymentType        CoreapiPaymentType          `json:"payment_type"`
	TransactionDetails midtrans.TransactionDetails `json:"transaction_details"`

	CreditCard   *CreditCardDetails   `json:"credit_card,omitempty"`
	BankTransfer *BankTransferDetails `json:"bank_transfer,omitempty"`

	CustomExpiry *CustomExpiry `json:"custom_expiry,omitempty"`
}

type CreditCardDetails struct {
	Bank string `json:"bank,omitempty"`
}

type BankTransferDetails struct {
	Bank     midtrans.Bank                  `json:"bank"`
	VaNumber string                         `json:"va_number,omitempty"`
	Permata  *PermataBankTransferDetail     `json:"permata,omitempty"`
	FreeText *BCABankTransferDetailFreeText `json:"free_text,omitempty"`
	Bca      *BcaBankTransferDetail         `json:"bca,omitempty"`
}

type PermataBankTransferDetail struct {
	RecipientName string `json:"recipient_name,omitempty"`
}

type BCABankTransferDetailFreeText struct {
	Inquiry []BCABankTransferLangDetail `json:"inquiry,omitempty"`
	Payment []BCABankTransferLangDetail `json:"payment,omitempty"`
}

type BCABankTransferLangDetail struct {
	LangID string `json:"id,omitempty"`
	LangEN string `json:"en,omitempty"`
}

type BcaBankTransferDetail struct {
	SubCompanyCode string `json:"sub_company_code,omitempty"`
}

type CustomExpiry struct {
	OrderTime      string `json:"order_time,omitempty"`
	ExpiryDuration int    `json:"expiry_duration,omitempty"`
	Unit           string `json:"unit,omitempty"`
}
