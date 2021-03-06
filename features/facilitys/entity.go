package facilitys

import "time"

type Core struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	InsertData(insert Core) (row int, err error)
	GetDataAll() (data []Core, err error)
	GetData(id int) (data Core, err error)
	UpdateData(id int, insert Core) (row int, err error)
	DeleteData(id int) (row int, err error)
}

type Data interface {
	InsertData(insert Core) (row int, err error)
	GetDataAll() (data []Core, err error)
	GetData(id int) (data Core, err error)
	UpdateData(id int, insert Core) (row int, err error)
	DeleteData(id int) (row int, err error)
}
