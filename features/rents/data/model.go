package data

import (
	"be9/mnroom/features/rents"
	"be9/mnroom/features/rooms"
	_rooms "be9/mnroom/features/rooms/data"
	"be9/mnroom/features/users"
	_users "be9/mnroom/features/users/data"

	"gorm.io/gorm"
)

type Rents struct {
	gorm.Model
	Date_start         string       `json:"date_start" form:"date_start"`
	Date_end           string       `json:"date_end" form:"date_end"`
	Bank               string       `json:"bank" form:"bank"`
	Total_rental_price int          `json:"total_rental_price" form:"total_rental_price"`
	Status             string       `json:"status" form:"status"`
	UserID             uint         `json:"user_id" form:"user_id"`
	RoomsID            uint         `json:"rooms_id" form:"rooms_id"`
	User               _users.User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rooms              _rooms.Rooms `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func toCoreList(data []Rents) []rents.Core {
	result := []rents.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *Rents) toCore() rents.Core {
	return rents.Core{
		ID:                 int(data.ID),
		Date_start:         data.Date_start,
		Date_end:           data.Date_end,
		Bank:               data.Bank,
		Total_rental_price: data.Total_rental_price,
		Status:             data.Status,
		User: users.Core{
			ID:       int(data.User.ID),
			Image:    data.User.Image,
			Username: data.User.Username,
			Email:    data.User.Email,
			Password: data.User.Password,
			Phone:    data.User.Phone,
			Address:  data.User.Address,
		},
		Room: rooms.Core{
			ID:          int(data.Rooms.ID),
			ImageRoom:   data.Rooms.ImageRoom,
			ImageLogo:   data.Rooms.ImageLogo,
			RoomName:    data.Rooms.RoomName,
			Capacity:    data.Rooms.Capacity,
			RentalPrice: data.Rooms.RentalPrice,
			City:        data.Rooms.City,
			Address:     data.Rooms.Address,
		},
	}
}

func fromCore(core rents.Core) Rents {
	return Rents{
		Date_start:         core.Date_start,
		Date_end:           core.Date_end,
		Bank:               core.Bank,
		Total_rental_price: core.Total_rental_price,
		Status:             core.Status,
		UserID:             uint(core.User.ID),
		RoomsID:            uint(core.Room.ID),
	}
}
