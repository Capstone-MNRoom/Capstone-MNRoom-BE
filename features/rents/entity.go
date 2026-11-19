package rents

import (
	"be9/mnroom/features/rooms"
	"be9/mnroom/features/users"
	"context"
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
	GetData(ctx context.Context, id int) (data int, err error)
	GetDataRentToken(ctx context.Context, idToken int, idRoom int) (row int, err error)
	GetDataRentUser(ctx context.Context, id int, start string, end string) (row int, err error)
	InsertData(ctx context.Context, insert Core) (row int, err error)
	GetDataRent(ctx context.Context, id int) (data []Core, err error)
	InsertDataPayment(ctx context.Context, insert CorePayments) (data CorePayments, err error)
	GetDataUser(ctx context.Context, idToken int) (data CoreUser, err error)
	GetDataRentUserHistory(ctx context.Context, idToken int) (data []Core, err error)
}

type Data interface {
	GetData(ctx context.Context, id int) (data int, err error)
	GetDataRentToken(ctx context.Context, idToken int, idRoom int) (row int, err error)
	GetDataRentUser(ctx context.Context, id int, start string, end string) (row int, err error)
	InsertData(ctx context.Context, insert Core) (row int, err error)
	GetDataRent(ctx context.Context, id int) (data []Core, err error)
	InsertDataPayment(ctx context.Context, insert CorePayments) (data CorePayments, err error)
	GetDataUser(ctx context.Context, idToken int) (data CoreUser, err error)
	GetDataRentUserHistory(ctx context.Context, idToken int) (data []Core, err error)
}
