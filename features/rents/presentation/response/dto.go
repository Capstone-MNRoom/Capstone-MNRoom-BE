package response

import (
	"be9/mnroom/features/rents"
	_rooms "be9/mnroom/features/rooms/presentation/response"
	_users "be9/mnroom/features/users/presentation/response"
)

type Rents struct {
	ID                 int    `json:"id"`
	Date_start         string `json:"date_start"`
	Date_end           string `json:"date_end"`
	Bank               string `json:"bank"`
	Total_rental_price int    `json:"total_rental_price"`
	Status             string `json:"status"`
	User               _users.User
	Rooms              _rooms.Rooms
}

func FromCoreList(data []rents.Core) []Rents {
	result := []Rents{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data rents.Core) Rents {
	return Rents{
		ID:                 data.ID,
		Date_start:         data.Date_start,
		Date_end:           data.Date_end,
		Bank:               data.Bank,
		Total_rental_price: data.Total_rental_price,
		Status:             data.Status,
		User: _users.User{
			ID:       data.User.ID,
			Image:    data.User.Image,
			Username: data.User.Username,
			Email:    data.User.Email,
			Phone:    data.User.Phone,
			Address:  data.User.Address,
		},
		Rooms: _rooms.Rooms{
			ID:             data.Room.ID,
			ImageRoom:      data.Room.ImageRoom,
			ImagePengelola: data.Room.ImagePengelola,
			Name:           data.Room.Name,
			Capacity:       data.Room.Capacity,
			RentalPrice:    data.Room.RentalPrice,
			City:           data.Room.City,
			Address:        data.Room.Address,
		},
	}
}
