package request

import (
	"be9/mnroom/features/categorys"
	"be9/mnroom/features/rooms"
)

type Rooms struct {
	ImageRoom      string `json:"image_room" validate:"required" form:"image_room"`
	ImagePengelola string `json:"image_pengelola" validate:"required" form:"image_pengelola"`
	Name           string `json:"name" validate:"required" form:"name"`
	Capacity       int    `json:"capacity" validate:"required,numeric" form:"capacity"`
	RentalPrice    int    `json:"rental_price" validate:"required,numeric" form:"rental_price"`
	Status         string `json:"status" validate:"required" form:"status"`
	Address        string `json:"address" validate:"required" form:"address"`
	City           string `json:"city" validate:"required" form:"city"`
	CategorysID    uint   `json:"categorys_id" validate:"required" form:"categorys_id"`
}

func ToCore(req Rooms) rooms.Core {
	return rooms.Core{
		ImageRoom:      req.ImageRoom,
		ImagePengelola: req.ImagePengelola,
		Name:           req.Name,
		Capacity:       req.Capacity,
		RentalPrice:    req.RentalPrice,
		Status:         req.Status,
		Address:        req.Address,
		City:           req.City,
		Categorys: categorys.Core{
			ID: int(req.CategorysID),
		},
	}
}
