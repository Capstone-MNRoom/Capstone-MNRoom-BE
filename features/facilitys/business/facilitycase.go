package business

import "be9/mnroom/features/facilitys"

type facilityCase struct {
	facilityData facilitys.Data
}

func NewFacilityBusiness(fctData facilitys.Data) facilitys.Business {
	return &facilityCase{
		facilityData: fctData,
	}
}

func (ctg *facilityCase) InsertData(insert facilitys.Core) (row int, err error) {
	row, err = ctg.facilityData.InsertData(insert)
	return row, err
}

func (ctg *facilityCase) GetDataAll() (data []facilitys.Core, err error) {
	data, err = ctg.facilityData.GetDataAll()
	return data, err
}

func (ctg *facilityCase) GetData(id int) (data facilitys.Core, err error) {
	data, err = ctg.facilityData.GetData(id)
	return data, err
}

func (ctg *facilityCase) UpdateData(id int, insert facilitys.Core) (row int, err error) {
	row, err = ctg.facilityData.UpdateData(id, insert)
	return row, err
}

func (ctg *facilityCase) DeleteData(id int) (row int, err error) {
	row, err = ctg.facilityData.DeleteData(id)
	return row, err
}
