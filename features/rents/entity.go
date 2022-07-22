package rents

import (
	"be9/mnroom/features/rooms"
	"be9/mnroom/features/users"
	"time"
)

type Core struct {
	ID               int
	DateStart        string
	DateEnd          string
	Bank             string
	TotalRentalPrice int
	Status           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	User             users.Core
	Room             rooms.Core
}

type CorePayments struct {
	ID                int
	TransactionID     string
	OrderID           int
	PaymentType       string
	BankTransfer      string
	GrossAmount       int
	VANumber          int
	TransactionStatus string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	User              users.Core
	Rents             Core
}

type Business interface {
	GetData(id int) (data int, err error)
	GetDataRentToken(idToken int) (row int, err error)
	GetDataRentUser(id int, start string, end string) (row int, err error)
	InsertData(insert Core) (row int, err error)
	GetDataRent(id int) (data []Core, err error)
	InsertDataPayment(insert CorePayments) (data CorePayments, err error)
}

type Data interface {
	GetData(id int) (data int, err error)
	GetDataRentToken(idToken int) (row int, err error)
	GetDataRentUser(id int, start string, end string) (row int, err error)
	InsertData(insert Core) (row int, err error)
	GetDataRent(id int) (data []Core, err error)
	InsertDataPayment(insert CorePayments) (data CorePayments, err error)
}
