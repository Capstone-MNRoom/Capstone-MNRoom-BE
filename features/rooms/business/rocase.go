package business

import "be9/mnroom/features/rooms"

type roomUseCase struct {
	roomData rooms.Data
}

func NewRoomBusiness(roData rooms.Data) rooms.Business {
	return &roomUseCase{
		roomData: roData,
	}
}

func (rocase *roomUseCase) InsertData(insert rooms.Core) (row int, err error) {
	row, err = rocase.roomData.InsertData(insert)
	return row, err
}

func (evcase *roomUseCase) GetAllData(limit int, offset int) (data []rooms.Core, err error) {
	data, err = evcase.roomData.GetAllData(limit, offset)
	return data, err
}

func (evcase *roomUseCase) GetData(id int) (data rooms.Core, err error) {
	data, err = evcase.roomData.GetData(id)
	return data, err
}

func (evcase *roomUseCase) DeleteData(id int) (row int, err error) {
	row, err = evcase.roomData.DeleteData(id)
	return row, err
}
