package response

import (
	"be9/mnroom/features/rooms"
	"time"
)

type Rooms struct {
	ID             int       `json:"id"`
	ImageRoom      string    `json:"image_room"`
	ImagePengelola string    `json:"image_pengelola"`
	RoomName       string    `json:"room_name"`
	Capacity       int       `json:"capacity"`
	HotelName      string    `json:"hotel_name"`
	RentalPrice    int       `json:"rental_price"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	CreatedAt      time.Time `json:"created_at"`
	User           User
	Categorys      Categorys
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Categorys struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

func FromCoreList(data []rooms.Core) []Rooms {
	result := []Rooms{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data rooms.Core) Rooms {
	return Rooms{
		ID:             data.ID,
		ImageRoom:      data.ImageRoom,
		ImagePengelola: data.ImagePengelola,
		RoomName:       data.RoomName,
		Capacity:       data.Capacity,
		HotelName:      data.HotelName,
		RentalPrice:    data.RentalPrice,
		Address:        data.Address,
		City:           data.City,
		CreatedAt:      data.CreatedAt,
		User: User{
			ID:       data.User.ID,
			Username: data.User.Username,
			Email:    data.User.Email,
		},
		Categorys: Categorys{
			ID:           data.Categorys.ID,
			CategoryName: data.Categorys.CategoryName,
		},
	}
}
