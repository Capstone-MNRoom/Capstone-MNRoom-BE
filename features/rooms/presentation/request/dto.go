package request

import (
	"be9/mnroom/features/categorys"
	"be9/mnroom/features/rooms"
)

type Rooms struct {
	ImageRoom      string `json:"image_room" validate:"required" form:"image_room"`
	ImagePengelola string `json:"image_pengelola" validate:"required" form:"image_pengelola"`
	RoomName       string `json:"room_name" form:"room_name"`
	Capacity       int    `json:"capacity" form:"capacity"`
	HotelName      string `json:"hotel_name" form:"hotel_name"`
	RentalPrice    int    `json:"rental_price" form:"rental_price"`
	Address        string `json:"address" form:"address"`
	Facilitys      []int  `json:"facilitys" form:"facilitys"`
	City           string `json:"city" form:"city"`
	CategorysID    uint   `json:"categorys_id" form:"categorys_id"`
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
