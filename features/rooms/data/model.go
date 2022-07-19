package data

import (
	"be9/mnroom/features/categorys"
	_categorys "be9/mnroom/features/categorys/data"
	_facilitys "be9/mnroom/features/facilitys/data"
	"be9/mnroom/features/rooms"
	"be9/mnroom/features/users"
	_users "be9/mnroom/features/users/data"

	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	ImageRoom      string               `json:"image_room" form:"image_room"`
	ImagePengelola string               `json:"image_pengelola" form:"image_pengelola"`
	RoomName       string               `json:"room_name" form:"room_name"`
	Capacity       int                  `json:"capacity" form:"capacity"`
	HotelName      string               `json:"hotel_name" form:"hotel_name"`
	RentalPrice    int                  `json:"rental_price" form:"rental_price"`
	Address        string               `json:"address" form:"address"`
	City           string               `json:"city" form:"city"`
	UserID         uint                 `json:"user_id" form:"user_id"`
	CategorysID    uint                 `json:"categorys_id" form:"categorys_id"`
	User           _users.User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Categorys      _categorys.Categorys `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type RoomFacilitys struct {
	gorm.Model
	UserID      uint `json:"users_id" form:"users_id"`
	RoomsID     uint `json:"rooms_id" form:"rooms_id"`
	FacilitysID uint `json:"facilitys_id" form:"facilitys_id"`
	User        _users.User
	Rooms       Rooms
	Facilitys   _facilitys.Facilitys
}

func toCoreList(data []Rooms) []rooms.Core {
	result := []rooms.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *Rooms) toCore() rooms.Core {
	return rooms.Core{
		ID:             int(data.ID),
		ImageRoom:      data.ImageRoom,
		ImagePengelola: data.ImagePengelola,
		RoomName:       data.RoomName,
		Capacity:       data.Capacity,
		HotelName:      data.HotelName,
		RentalPrice:    data.RentalPrice,
		Address:        data.Address,
		City:           data.City,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
		User: users.Core{
			ID:       int(data.User.ID),
			Username: data.User.Username,
		},
		Categorys: categorys.Core{
			ID:           int(data.Categorys.ID),
			CategoryName: data.Categorys.CategoryName,
		},
	}
}

func fromCore(core rooms.Core) Rooms {
	return Rooms{
		ImageRoom:      core.ImageRoom,
		ImagePengelola: core.ImagePengelola,
		RoomName:       core.RoomName,
		Capacity:       core.Capacity,
		HotelName:      core.HotelName,
		RentalPrice:    core.RentalPrice,
		Address:        core.Address,
		City:           core.City,
		UserID:         uint(core.User.ID),
		CategorysID:    uint(core.Categorys.ID),
	}
}
