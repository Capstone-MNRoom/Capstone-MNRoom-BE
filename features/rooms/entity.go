package rooms

import (
	"be9/mnroom/features/categorys"
	"be9/mnroom/features/users"
	"time"
)

type Core struct {
	ID             int
	ImageRoom      string
	ImagePengelola string
	Name           string
	Capacity       int
	RentalPrice    int
	Address        string
	City           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           users.Core
	Categorys      categorys.Core
}

type Business interface {
	InsertData(insert Core) (row int, err error)
	GetDataAll() (data []Core, err error)
	GetData(id int) (data Core, err error)
	UpdateData(id int, insert Core) (row int, err error)
	DeleteData(id int) (row int, err error)
	GetToken(id int, idToken int) (data Core, err error)
	GetDataAllUserRoom(idToken int) (data []Core, err error)
}

type Data interface {
	InsertData(insert Core) (row int, err error)
	GetDataAll() (data []Core, err error)
	GetData(id int) (data Core, err error)
	UpdateData(id int, insert Core) (row int, err error)
	DeleteData(id int) (row int, err error)
	GetToken(id int, idToken int) (data Core, err error)
	GetDataAllUserRoom(idToken int) (data []Core, err error)
}
