package business

import "be9/mnroom/features/categorys"

type categoryCase struct {
	categoryData categorys.Data
}

func NewCategoryBusiness(ctgData categorys.Data) categorys.Business {
	return &categoryCase{
		categoryData: ctgData,
	}
}

func (ctg *categoryCase) InsertData(insert categorys.Core) (row int, err error) {
	row, err = ctg.categoryData.InsertData(insert)
	return row, err
}

func (ctg *categoryCase) GetDataAll() (data []categorys.Core, err error) {
	data, err = ctg.categoryData.GetDataAll()
	return data, err
}

func (ctg *categoryCase) GetData(id int) (data categorys.Core, err error) {
	data, err = ctg.categoryData.GetData(id)
	return data, err
}

func (ctg *categoryCase) UpdateData(id int, insert categorys.Core) (row int, err error) {
	row, err = ctg.categoryData.UpdateData(id, insert)
	return row, err
}

func (ctg *categoryCase) DeleteData(id int) (row int, err error) {
	row, err = ctg.categoryData.DeleteData(id)
	return row, err
}
