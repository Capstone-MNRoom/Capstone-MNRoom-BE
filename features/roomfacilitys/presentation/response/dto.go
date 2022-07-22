package response

import (
	"be9/mnroom/features/roomfacilitys"
	"time"
)

type RoomFacilitys struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	User      User
	Rooms     Rooms
	Facilitys Facilitys
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
	Capacity  int    `json:"capacity"`
}

type Facilitys struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCoreList(data []roomfacilitys.Core) []RoomFacilitys {
	result := []RoomFacilitys{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data roomfacilitys.Core) RoomFacilitys {
	return RoomFacilitys{
		ID:        data.ID,
		CreatedAt: data.CreatedAt,
		User: User{
			ID:       data.User.ID,
			Username: data.User.Username,
			Email:    data.User.Email,
		},
		Rooms: Rooms{
			ID:        data.Rooms.ID,
			RoomName:  data.Rooms.RoomName,
			HotelName: data.Rooms.HotelName,
			Capacity:  data.Rooms.Capacity,
		},
		Facilitys: Facilitys{
			ID:   data.Facilitys.ID,
			Name: data.Facilitys.Name,
		},
	}
}
