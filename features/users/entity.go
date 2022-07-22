package users

import "time"

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
	InsertData(insert Core) (row int, err error)
	GetAllData() (data []Core, err error)
	GetData(id int) (data Core, err error)
	DeleteData(id int) (row int, err error)
	UpdateData(id int, insert Core) (row int, err error)
}

type Data interface {
	InsertData(insert Core) (row int, err error)
	GetAllData() (data []Core, err error)
	GetData(id int) (data Core, err error)
	DeleteData(id int) (row int, err error)
	UpdateData(id int, insert Core) (row int, err error)
}
