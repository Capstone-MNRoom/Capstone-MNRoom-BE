package midtrans

type CoreapiPaymentType string
type SubscriptionPaymentType = CoreapiPaymentType

const (
	PaymentTypeBankTransfer CoreapiPaymentType = "bank_transfer"
)
