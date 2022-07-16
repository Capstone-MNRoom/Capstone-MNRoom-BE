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
	Status         string
	Address        string
	City           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           users.Core
	Categorys      categorys.Core
}

type Business interface {
}

type Data interface {
}
