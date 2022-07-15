package business

import (
	"be9/mnroom/features/users"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(insert.Password), bcrypt.DefaultCost)
	if errHash != nil {
		return 0, fmt.Errorf("failed to generate password")
	}
	insert.Password = string(hashedPassword)
	row, err = usecase.userData.InsertData(insert)
	return row, err
}

func (usecase *userUsecase) GetAllData(limit int, offset int) (data []users.Core, err error) {
	data, err = usecase.userData.GetAllData(limit, offset)
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
