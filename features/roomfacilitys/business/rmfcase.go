package business

import (
	"be9/mnroom/features/roomfacilitys"
	"context"
)

type roomfacilityCase struct {
	roomfacilityData roomfacilitys.Data
}

func NewRoomFacilityBusiness(rmfData roomfacilitys.Data) roomfacilitys.Business {
	return &roomfacilityCase{
		roomfacilityData: rmfData,
	}
}

func (rmf *roomfacilityCase) GetData(ctx context.Context, id int) (data []roomfacilitys.Core, err error) {
	data, err = rmf.roomfacilityData.GetData(ctx, id)
	return data, err
}

func (rmf *roomfacilityCase) GetDataRow(ctx context.Context, id int) (row int, err error) {
	row, err = rmf.roomfacilityData.GetDataRow(ctx, id)
	return row, err
}
