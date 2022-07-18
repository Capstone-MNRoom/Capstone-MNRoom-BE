package response

import (
	"be9/mnroom/features/rents"
	"time"
)

type Rents struct {
	ID               int       `json:"id"`
	DateStart        string    `json:"date_start"`
	DateEnd          string    `json:"date_end"`
	Bank             string    `json:"bank"`
	TotalRentalPrice int       `json:"total_rental_price"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	User             User
	Rooms            Rooms
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Rooms struct {
	ID        int    `json:"id"`
	RoomName  string `json:"room_name"`
	HotelName string `json:"hotel_name"`
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
		ID:               data.ID,
		DateStart:        data.DateStart,
		DateEnd:          data.DateEnd,
		Bank:             data.Bank,
		TotalRentalPrice: data.TotalRentalPrice,
		Status:           data.Status,
		User: User{
			ID:       data.User.ID,
			Username: data.User.Username,
		},
		Rooms: Rooms{
			ID:        data.Room.ID,
			RoomName:  data.Room.RoomName,
			HotelName: data.Room.HotelName,
		},
	}
}
