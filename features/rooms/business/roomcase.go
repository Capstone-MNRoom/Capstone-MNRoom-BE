package business

import "be9/mnroom/features/rooms"

type roomCase struct {
	roomData rooms.Data
}

func NewRoomBusiness(rmsData rooms.Data) rooms.Business {
	return &roomCase{
		roomData: rmsData,
	}
}
