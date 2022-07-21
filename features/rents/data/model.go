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
	DateStart        string       `json:"date_start" form:"date_start"`
	DateEnd          string       `json:"date_end" form:"date_end"`
	Bank             string       `json:"bank" form:"bank"`
	TotalRentalPrice int          `json:"total_rental_price" form:"total_rental_price"`
	Status           string       `json:"status" form:"status"`
	UserID           uint         `json:"user_id" form:"user_id"`
	RoomsID          uint         `json:"rooms_id" form:"rooms_id"`
	User             _users.User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rooms            _rooms.Rooms `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
		ID:               int(data.ID),
		DateStart:        data.DateStart,
		DateEnd:          data.DateEnd,
		Bank:             data.Bank,
		TotalRentalPrice: data.TotalRentalPrice,
		Status:           data.Status,
		User: users.Core{
			ID:       int(data.User.ID),
			Username: data.User.Username,
		},
		Room: rooms.Core{
			ID:        int(data.Rooms.ID),
			RoomName:  data.Rooms.RoomName,
			HotelName: data.Rooms.HotelName,
			ImageRoom: data.Rooms.ImageRoom,
		},
	}
}

func fromCore(core rents.Core) Rents {
	return Rents{
		DateStart:        core.DateStart,
		DateEnd:          core.DateEnd,
		Bank:             core.Bank,
		TotalRentalPrice: core.TotalRentalPrice,
		Status:           core.Status,
		UserID:           uint(core.User.ID),
		RoomsID:          uint(core.Room.ID),
	}
}
