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

type User struct {
	ID       int
	Username string
	Email    string
}

type Rents struct {
	ID               int
	DateStart        string
	DateEnd          string
	Bank             string
	TotalRentalPrice int
	Status           string
}

type Business interface {
	GetAllData(idToken int) (data []Core, err error)
}

type Data interface {
	GetAllData(idToken int) (data []Core, err error)
}
