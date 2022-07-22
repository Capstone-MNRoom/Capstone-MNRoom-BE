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

func (rm *roomCase) InsertData(insert rooms.Core) (data rooms.Core, err error) {
	data, err = rm.roomData.InsertData(insert)
	return data, err
}

func (rm *roomCase) InsertDataRoomFacilitys(insert rooms.CoreRoomFacilitys) (row int, err error) {
	row, err = rm.roomData.InsertDataRoomFacilitys(insert)
	return row, err
}

func (rm *roomCase) GetDataAll() (data []rooms.Core, err error) {
	data, err = rm.roomData.GetDataAll()
	return data, err
}

func (rm *roomCase) GetData(id int) (data rooms.Core, err error) {
	data, err = rm.roomData.GetData(id)
	return data, err
}

func (rm *roomCase) GetDataIDRoom(id int) (row int, err error) {
	row, err = rm.roomData.GetDataIDRoom(id)
	return row, err
}

func (rm *roomCase) UpdateData(id int, insert rooms.Core) (row int, err error) {
	row, err = rm.roomData.UpdateData(id, insert)
	return row, err
}

func (rm *roomCase) DeleteData(id int) (row int, err error) {
	row, err = rm.roomData.DeleteData(id)
	return row, err
}

func (rm *roomCase) GetToken(id int, idToken int) (data rooms.Core, err error) {
	data, err = rm.roomData.GetToken(id, idToken)
	return data, err
}

func (rm *roomCase) GetDataAllUserRoom(idToken int) (data []rooms.Core, err error) {
	data, err = rm.roomData.GetDataAllUserRoom(idToken)
	return data, err
}
