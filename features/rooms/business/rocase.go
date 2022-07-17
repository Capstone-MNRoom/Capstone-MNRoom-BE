package business

import (
	"be9/mnroom/features/rooms"
)

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

func (rocase *roomUseCase) GetAllData(limit int, offset int) (data []rooms.Core, err error) {
	data, err = rocase.roomData.GetAllData(limit, offset)
	return data, err
}

func (rocase *roomUseCase) GetData(id int) (data rooms.Core, err error) {
	data, err = rocase.roomData.GetData(id)
	return data, err
}

func (rocase *roomUseCase) DeleteData(id int) (row int, err error) {
	row, err = rocase.roomData.DeleteData(id)
	return row, err
}
func (rocase *roomUseCase) GetToken(id int, idToken int) (data rooms.Core, err error) {
	data, err = rocase.roomData.GetToken(id, idToken)
	return data, err
}

func (evcase *roomUseCase) UpdatedData(id int, insert rooms.Core) (row int, err error) {
	row, err = evcase.roomData.UpdatedData(id, insert)
	return row, err
}