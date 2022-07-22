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
	Email    string `json:"email"`
}

type Rooms struct {
	ID        int    `json:"id"`
	RoomName  string `json:"room_name"`
	HotelName string `json:"hotel_name"`
	ImageRoom string `json:"image_room"`
	Capacity  int    `json:"capacity"`
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
		TotalRentalPrice: data.TotalRentalPrice,
		User: User{
			ID:       data.User.ID,
			Username: data.User.Username,
			Email:    data.User.Email,
		},
		Rooms: Rooms{
			ID:        data.Room.ID,
			RoomName:  data.Room.RoomName,
			HotelName: data.Room.HotelName,
			ImageRoom: data.Room.ImageRoom,
			Capacity:  data.Room.Capacity,
		},
	}
}
