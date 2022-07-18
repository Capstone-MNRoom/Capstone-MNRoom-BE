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

type Business interface {
}

type Data interface {
}
