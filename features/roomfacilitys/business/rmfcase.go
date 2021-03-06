package business

import "be9/mnroom/features/roomfacilitys"

type roomfacilityCase struct {
	roomfacilityData roomfacilitys.Data
}

func NewRoomFacilityBusiness(rmfData roomfacilitys.Data) roomfacilitys.Business {
	return &roomfacilityCase{
		roomfacilityData: rmfData,
	}
}

func (rmf *roomfacilityCase) GetData(id int) (data []roomfacilitys.Core, err error) {
	data, err = rmf.roomfacilityData.GetData(id)
	return data, err
}

func (rmf *roomfacilityCase) GetDataRow(id int) (row int, err error) {
	row, err = rmf.roomfacilityData.GetDataRow(id)
	return row, err
}
