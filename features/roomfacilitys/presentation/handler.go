package presentation

import (
	"be9/mnroom/features/roomfacilitys"
)

type RoomFacilityHandler struct {
	roomfacilitysBusiness roomfacilitys.Business
}

func NewRoomFacilitysHandler(business roomfacilitys.Business) *RoomFacilityHandler {
	return &RoomFacilityHandler{
		roomfacilitysBusiness: business,
	}
}
