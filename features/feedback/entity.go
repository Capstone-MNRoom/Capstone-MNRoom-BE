package feedback

import (
	"time"
)

type Core struct {
	ID        int
	Rating    int
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Rents     Rents
}

type User struct {
	ID       int
	Username string
	Email    string
}

type Rents struct {
	ID                 int
	Date_start         string
	Date_end           string
	Bank               string
	Total_rental_price int
	Status             string
	Room               Rooms
}

type Rooms struct {
	ID int
}

type Business interface {
	GetFeedbackByRoom(id int) (data []Core, err error)
	InsertFeedback(insert Core) (row int, err error)
	GetDataRoom(id int) (data int, err error)
	GetDataRent(id int) (data int, err error)
	GetDataRentUser(idToken int) (data int, err error)
}

type Data interface {
	GetFeedbackByRoom(id int) (data []Core, err error)
	InsertFeedback(insert Core) (row int, err error)
	GetDataRoom(id int) (data int, err error)
	GetDataRent(id int) (data int, err error)
	GetDataRentUser(idToken int) (data int, err error)
}
