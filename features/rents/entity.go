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
	OrderID           string
	PaymentType       string
	BankTransfer      string
	GrossAmount       int
	VANumber          string
	TransactionStatus string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	User              users.Core
	Rents             Core
}

type CoreUser struct {
	ID        int
	Image     string
	Username  string
	Email     string
	Password  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	GetData(id int) (data int, err error)
	GetDataRentToken(idToken int, idRoom int) (row int, err error)
	GetDataRentUser(id int, start string, end string) (row int, err error)
	InsertData(insert Core) (row int, err error)
	GetDataRent(id int) (data []Core, err error)
	InsertDataPayment(insert CorePayments) (data CorePayments, err error)
	GetDataUser(idToken int) (data CoreUser, err error)
	GetDataRentUserHistory(idToken int) (data []Core, err error)
}

type Data interface {
	GetData(id int) (data int, err error)
	GetDataRentToken(idToken int, idRoom int) (row int, err error)
	GetDataRentUser(id int, start string, end string) (row int, err error)
	InsertData(insert Core) (row int, err error)
	GetDataRent(id int) (data []Core, err error)
	InsertDataPayment(insert CorePayments) (data CorePayments, err error)
	GetDataUser(idToken int) (data CoreUser, err error)
	GetDataRentUserHistory(idToken int) (data []Core, err error)
}
