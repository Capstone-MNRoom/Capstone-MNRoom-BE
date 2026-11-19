package business

import (
	"be9/mnroom/features/facilitys"
	"context"
)

type facilityCase struct {
	facilityData facilitys.Data
}

func NewFacilityBusiness(fctData facilitys.Data) facilitys.Business {
	return &facilityCase{
		facilityData: fctData,
	}
}

func (ctg *facilityCase) InsertData(ctx context.Context, insert facilitys.Core) (row int, err error) {
	row, err = ctg.facilityData.InsertData(ctx, insert)
	return row, err
}

func (ctg *facilityCase) GetDataAll(ctx context.Context) (data []facilitys.Core, err error) {
	data, err = ctg.facilityData.GetDataAll(ctx)
	return data, err
}

func (ctg *facilityCase) GetData(ctx context.Context, id int) (data facilitys.Core, err error) {
	data, err = ctg.facilityData.GetData(ctx, id)
	return data, err
}

func (ctg *facilityCase) UpdateData(ctx context.Context, id int, insert facilitys.Core) (row int, err error) {
	row, err = ctg.facilityData.UpdateData(ctx, id, insert)
	return row, err
}

func (ctg *facilityCase) DeleteData(ctx context.Context, id int) (row int, err error) {
	row, err = ctg.facilityData.DeleteData(ctx, id)
	return row, err
}
