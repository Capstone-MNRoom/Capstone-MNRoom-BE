package rooms

import (
	"be9/mnroom/features/categorys"
	"be9/mnroom/features/facilitys"
	"be9/mnroom/features/users"
	"time"
)

type Core struct {
	ID          int
	ImageRoom   string
	ImageLogo   string
	RoomName    string
	Capacity    int
	RentalPrice int
	Status      string
	City        string
	Address     string
	Deskripsi   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        users.Core
	Categorys   categorys.Core
	Facilitys   facilitys.Core
}

type Business interface {
	InsertData(insert Core) (row int, err error)
	// GetAllData(limit int, offset int) (data []Core, err error)
	// GetData(id int) (data Core, err error)
	// DeleteData(id int) (row int, err error)
	// GetToken(id int, idToken int) (data Core, err error)
	// UpdatedData(id int, insert Core) (row int, err error)

}

type Data interface {
	InsertData(insert Core) (row int, err error)
	// GetAllData(limit int, offset int) (data []Core, err error)
	// GetData(id int) (data Core, err error)
	// DeleteData(id int) (row int, err error)
	// GetToken(id int, idToken int) (data Core, err error)
	// UpdatedData(id int, insert Core) (row int, err error)

}
