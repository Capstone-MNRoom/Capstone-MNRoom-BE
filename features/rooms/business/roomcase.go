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

func (rm *roomCase) InsertData(insert rooms.Core) (row int, err error) {
	row, err = rm.roomData.InsertData(insert)
	return row, err
}
