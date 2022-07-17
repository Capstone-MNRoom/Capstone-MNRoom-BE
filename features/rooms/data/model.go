package data

import (
	"be9/mnroom/features/categorys"
	_categorys "be9/mnroom/features/categorys/data"
	"be9/mnroom/features/rooms"
	"be9/mnroom/features/users"
	_users "be9/mnroom/features/users/data"

	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	ImageRoom   string               `json:"image_room" form:"image_room"`
	ImageLogo   string               `json:"image_logo" form:"image_logo"`
	RoomName    string               `json:"room_name" form:"room_name"`
	Capacity    int                  `json:"capacity" form:"capacity"`
	RentalPrice int                  `json:"rental_price" form:"rental_price"`
	City        string               `json:"city" form:"city"`
	Address     string               `json:"address" form:"address"`
	UserID      uint                 `json:"user_id" form:"user_id"`
	CategoryID  uint                 `json:"category_id" form:"category_id"`
	User        _users.User          `gorm:"foreignKey:UserID;"`
	Categorys   _categorys.Categorys `gorm:"foreignKey:CategoryID;"`
	// User        users.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
		ID:          int(data.ID),
		ImageRoom:   data.ImageRoom,
		RoomName:    data.RoomName,
		Capacity:    data.Capacity,
		RentalPrice: data.RentalPrice,
		City:        data.City,
		Address:     data.Address,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		User: users.Core{
			ID:       int(data.User.ID),
			Image:    data.User.Image,
			Username: data.User.Username,
			Email:    data.User.Email,
			Phone:    data.User.Phone,
			Address:  data.User.Address,
		},
		Categorys: categorys.Core{
			ID:           int(data.Categorys.ID),
			CategoryName: data.Categorys.CategoryName,
		},
	}
}

func fromCore(core rooms.Core) Rooms {
	return Rooms{
		ImageRoom:   core.ImageRoom,
		ImageLogo:   core.ImageLogo,
		RoomName:    core.RoomName,
		Capacity:    core.Capacity,
		RentalPrice: core.RentalPrice,
		City:        core.City,
		Address:     core.Address,
		UserID:      uint(core.User.ID),
		CategoryID:  uint(core.Categorys.ID),
		// User:        uint(core.User.ID),
		// Categorys:   uint(core.Categorys.ID),

	}
}
