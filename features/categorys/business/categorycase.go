package business

import (
	"be9/mnroom/features/categorys"
	"context"
)

type categoryCase struct {
	categoryData categorys.Data
}

func NewCategoryBusiness(ctgData categorys.Data) categorys.Business {
	return &categoryCase{
		categoryData: ctgData,
	}
}

func (ctg *categoryCase) InsertData(ctx context.Context, insert categorys.Core) (row int, err error) {
	row, err = ctg.categoryData.InsertData(ctx, insert)
	return row, err
}

func (ctg *categoryCase) GetDataAll(ctx context.Context) (data []categorys.Core, err error) {
	data, err = ctg.categoryData.GetDataAll(ctx)
	return data, err
}

func (ctg *categoryCase) GetData(ctx context.Context, id int) (data categorys.Core, err error) {
	data, err = ctg.categoryData.GetData(ctx, id)
	return data, err
}

func (ctg *categoryCase) UpdateData(ctx context.Context, id int, insert categorys.Core) (row int, err error) {
	row, err = ctg.categoryData.UpdateData(ctx, id, insert)
	return row, err
}

func (ctg *categoryCase) DeleteData(ctx context.Context, id int) (row int, err error) {
	row, err = ctg.categoryData.DeleteData(ctx, id)
	return row, err
}
