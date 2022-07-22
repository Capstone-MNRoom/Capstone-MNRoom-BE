package business

import (
	"be9/mnroom/features/users"
)

type userUsecase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUsecase{
		userData: usrData,
	}
}

func (usecase *userUsecase) InsertData(insert users.Core) (row int, err error) {
	row, err = usecase.userData.InsertData(insert)
	return row, err
}

func (usecase *userUsecase) GetAllData() (data []users.Core, err error) {
	data, err = usecase.userData.GetAllData()
	return data, err
}

func (usecase *userUsecase) GetData(id int) (data users.Core, err error) {
	data, err = usecase.userData.GetData(id)
	return data, err
}

func (usecase *userUsecase) DeleteData(id int) (row int, err error) {
	row, err = usecase.userData.DeleteData(id)
	return row, err
}

func (usecase *userUsecase) UpdateData(id int, insert users.Core) (row int, err error) {
	row, err = usecase.userData.UpdateData(id, insert)
	return row, err
}
