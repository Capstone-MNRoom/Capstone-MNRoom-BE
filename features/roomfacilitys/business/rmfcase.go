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
