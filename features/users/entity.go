package users

import (
	"context"
	"time"
)

type Core struct {
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
	InsertData(ctx context.Context, insert Core) (row int, err error)
	GetAllData(ctx context.Context) (data []Core, err error)
	GetData(ctx context.Context, id int) (data Core, err error)
	DeleteData(ctx context.Context, id int) (row int, err error)
	UpdateData(ctx context.Context, id int, insert Core) (row int, err error)
}

type Data interface {
	InsertData(ctx context.Context, insert Core) (row int, err error)
	GetAllData(ctx context.Context) (data []Core, err error)
	GetData(ctx context.Context, id int) (data Core, err error)
	DeleteData(ctx context.Context, id int) (row int, err error)
	UpdateData(ctx context.Context, id int, insert Core) (row int, err error)
}
