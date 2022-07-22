package payments

import (
	"be9/mnroom/features/rents"
	"be9/mnroom/features/users"
	"time"
)

type Core struct {
	ID                int
	TransactionID     string
	OrderID           string
	PaymentType       string
	BankTransfer      string
	GrossAmount       int
	VANumber          string
	TransactionStatus string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	User              users.Core
	Rents             rents.Core
}

type Business interface {
}

type Data interface {
}
