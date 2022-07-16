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
}

type Data interface {
	InsertData(insert Core) (row int, err error)
	GetDataAll() (data []Core, err error)
}
