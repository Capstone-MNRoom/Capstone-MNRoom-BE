package request

import (
	"be9/mnroom/features/facilitys"
	"be9/mnroom/features/roomfacilitys"
	"be9/mnroom/features/rooms"
)

type RoomFacilitys struct {
	RoomsID     uint `json:"rooms_id" form:"rooms_id"`
	FacilitysID uint `json:"facilitys_id" form:"facilitys_id"`
}

func ToCore(req RoomFacilitys) roomfacilitys.Core {
	return roomfacilitys.Core{
		Rooms: rooms.Core{
			ID: int(req.RoomsID),
		},
		Facilitys: facilitys.Core{
			ID: int(req.FacilitysID),
		},
	}
}
