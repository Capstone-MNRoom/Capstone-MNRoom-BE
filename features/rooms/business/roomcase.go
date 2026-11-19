package business

import (
	"be9/mnroom/features/rooms"
	"context"
)

type roomCase struct {
	roomData rooms.Data
}

func NewRoomBusiness(rmsData rooms.Data) rooms.Business {
	return &roomCase{
		roomData: rmsData,
	}
}

func (rm *roomCase) InsertData(ctx context.Context, insert rooms.Core) (data rooms.Core, err error) {
	data, err = rm.roomData.InsertData(ctx, insert)
	return data, err
}

func (rm *roomCase) InsertDataRoomFacilitys(ctx context.Context, insert rooms.CoreRoomFacilitys) (row int, err error) {
	row, err = rm.roomData.InsertDataRoomFacilitys(ctx, insert)
	return row, err
}

func (rm *roomCase) GetDataAll(ctx context.Context, page int) (data []rooms.Core, err error) {
	limit := 12
	page = limit * (page - 1)
	data, err = rm.roomData.GetDataAll(ctx, page)
	return data, err
}

func (rm *roomCase) GetData(ctx context.Context, id int) (data rooms.Core, err error) {
	data, err = rm.roomData.GetData(ctx, id)
	return data, err
}

func (rm *roomCase) GetDataIDRoom(ctx context.Context, id int) (row int, err error) {
	row, err = rm.roomData.GetDataIDRoom(ctx, id)
	return row, err
}

func (rm *roomCase) UpdateData(ctx context.Context, id int, insert rooms.Core) (row int, err error) {
	row, err = rm.roomData.UpdateData(ctx, id, insert)
	return row, err
}

func (rm *roomCase) DeleteData(ctx context.Context, id int) (row int, err error) {
	row, err = rm.roomData.DeleteData(ctx, id)
	return row, err
}

func (rm *roomCase) GetToken(ctx context.Context, id int, idToken int) (data rooms.Core, err error) {
	data, err = rm.roomData.GetToken(ctx, id, idToken)
	return data, err
}

func (rm *roomCase) GetDataAllUserRoom(ctx context.Context, idToken int) (data []rooms.Core, err error) {
	data, err = rm.roomData.GetDataAllUserRoom(ctx, idToken)
	return data, err
}

func (rm *roomCase) GetDataByCategory(ctx context.Context, id int) (data []rooms.Core, err error) {
	data, err = rm.roomData.GetDataByCategory(ctx, id)
	return data, err
}
