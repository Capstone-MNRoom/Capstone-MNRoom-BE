package rents

import (
	"be9/mnroom/features/rooms"
	"be9/mnroom/features/users"
	"time"
)

type Core struct {
	ID                 int
	Date_start         string
	Date_end           string
	Bank               string
	Total_rental_price int
	Status             string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	User               users.Core
	Room               rooms.Core
}

type Business interface {
	InsertData(insert Core) (row int, err error)
}

type Data interface {
	InsertData(insert Core) (row int, err error)
}
