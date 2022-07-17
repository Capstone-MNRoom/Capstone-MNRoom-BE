package request

import (
	"be9/mnroom/features/categorys"
	"be9/mnroom/features/rooms"
)

type Rooms struct {
	ImageRoom   string `json:"image_room" form:"image_room"`
	ImageLogo   string `json:"image_logo" form:"image_logo"`
	RoomName    string `json:"room_name" form:"room_name"`
	Capacity    int    `json:"capacity" form:"capacity"`
	RentalPrice int    `json:"rental_price" form:"rental_price"`
	City        string `json:"city" form:"city"`
	Address     string `json:"address" form:"address"`
	CategorysID uint   `json:"categorys_id" validate:"required" form:"categorys_id"`
	// FacilityID  uint   `json:"facility_id" validate:"required" form:"facility_id"`
}

func ToCore(req Rooms) rooms.Core {
	return rooms.Core{
		ImageRoom:   req.ImageRoom,
		ImageLogo:   req.ImageLogo,
		RoomName:    req.RoomName,
		Capacity:    req.Capacity,
		RentalPrice: req.RentalPrice,
		City:        req.City,
		Address:     req.Address,
		Categorys: categorys.Core{
			ID: int(req.CategorysID),
		},
	}
}
