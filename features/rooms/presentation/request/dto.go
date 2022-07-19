package request

import (
	"be9/mnroom/features/categorys"
	"be9/mnroom/features/rooms"
)

type Rooms struct {
	ImageRoom      string `json:"image_room" validate:"required" form:"image_room"`
	ImagePengelola string `json:"image_pengelola" validate:"required" form:"image_pengelola"`
	RoomName       string `json:"room_name" validate:"required" form:"room_name"`
	Capacity       int    `json:"capacity" validate:"required,numeric" form:"capacity"`
	HotelName      string `json:"hotel_name" validate:"required" form:"hotel_name"`
	RentalPrice    int    `json:"rental_price" validate:"required,numeric" form:"rental_price"`
	Address        string `json:"address" validate:"required" form:"address"`
	Facilitys      []int  `json:"facilitys" form:"facilitys"`
	City           string `json:"city" validate:"required" form:"city"`
	CategorysID    uint   `json:"categorys_id" validate:"required" form:"categorys_id"`
}

func ToCore(req Rooms) rooms.Core {
	return rooms.Core{
		ImageRoom:      req.ImageRoom,
		ImagePengelola: req.ImagePengelola,
		RoomName:       req.RoomName,
		Capacity:       req.Capacity,
		HotelName:      req.HotelName,
		RentalPrice:    req.RentalPrice,
		Address:        req.Address,
		City:           req.City,
		Facilitys:      req.Facilitys,
		Categorys: categorys.Core{
			ID: int(req.CategorysID),
		},
	}
}
