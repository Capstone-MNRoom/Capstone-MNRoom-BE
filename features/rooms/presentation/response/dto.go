package response

import (
	_categorys "be9/mnroom/features/categorys/presentation/response"
	_facilitys "be9/mnroom/features/facilitys/presentation/response"
	"be9/mnroom/features/rooms"
	_users "be9/mnroom/features/users/presentation/response"
	"time"
)

type Rooms struct {
	ID          int       `json:"id"`
	ImageRoom   string    `json:"image_room"`
	ImageLogo   string    `json:"image_logo"`
	RoomName    string    `json:"room_name"`
	Capacity    int       `json:"capacity"`
	RentalPrice int       `json:"rental_price"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	Status      string    `json:"status"`
	Deskripsi   string    `json:"deskripsi"`
	CreatedAt   time.Time `json:"created_at"`
	Categorys   _categorys.Categorys
	Facility    _facilitys.Facilitys
	User        _users.User
}

func FromCoreList(data []rooms.Core) []Rooms {
	resuly := []Rooms{}
	for key := range data {
		resuly = append(resuly, FromCore(data[key]))
	}
	return resuly
}

func FromCore(data rooms.Core) Rooms {
	return Rooms{
		ID:          data.ID,
		ImageRoom:   data.ImageRoom,
		RoomName:    data.RoomName,
		Capacity:    data.Capacity,
		RentalPrice: data.RentalPrice,
		City:        data.City,
		Address:     data.Address,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
		User: _users.User{
			ID:       int(data.User.ID),
			Image:    data.User.Image,
			Username: data.User.Username,
			Email:    data.User.Email,
			Phone:    data.User.Phone,
			Address:  data.User.Address,
		},
		Categorys: _categorys.Categorys{
			ID:           data.Categorys.ID,
			CategoryName: data.Categorys.CategoryName,
		},
		Facility: _facilitys.Facilitys{
			ID:   data.Facilitys.ID,
			Name: data.Facilitys.Name,
		},
	}
}
