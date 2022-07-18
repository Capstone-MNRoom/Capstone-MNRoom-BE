package data

import (
	"be9/mnroom/features/facilitys"
	_facilitys "be9/mnroom/features/facilitys/data"
	"be9/mnroom/features/roomfacilitys"
	"be9/mnroom/features/rooms"
	_rooms "be9/mnroom/features/rooms/data"
	"be9/mnroom/features/users"
	_users "be9/mnroom/features/users/data"

	"gorm.io/gorm"
)

type RoomFacilitys struct {
	gorm.Model
	UserID      uint `json:"users_id" form:"users_id"`
	RoomsID     uint `json:"rooms_id" form:"rooms_id"`
	FacilitysID uint `json:"facilitys_id" form:"facilitys_id"`
	User        _users.User
	Rooms       _rooms.Rooms
	Facilitys   _facilitys.Facilitys
}

func toCoreList(data []RoomFacilitys) []roomfacilitys.Core {
	result := []roomfacilitys.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *RoomFacilitys) toCore() roomfacilitys.Core {
	return roomfacilitys.Core{
		ID: int(data.ID),
		User: users.Core{
			ID:       int(data.User.ID),
			Username: data.User.Username,
		},
		Rooms: rooms.Core{
			ID:   int(data.Rooms.ID),
			Name: data.Rooms.Name,
		},
		Facilitys: facilitys.Core{
			ID:   int(data.Facilitys.ID),
			Name: data.Facilitys.Name,
		},
	}
}

func fromCore(core roomfacilitys.Core) RoomFacilitys {
	return RoomFacilitys{
		RoomsID:     uint(core.Rooms.ID),
		FacilitysID: uint(core.Facilitys.ID),
	}
}
