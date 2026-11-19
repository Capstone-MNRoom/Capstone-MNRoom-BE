package business

import (
	"be9/mnroom/features/users"
	"context"
)

type userUsecase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUsecase{
		userData: usrData,
	}
}

func (usecase *userUsecase) InsertData(ctx context.Context, insert users.Core) (row int, err error) {
	row, err = usecase.userData.InsertData(ctx, insert)
	return row, err
}

func (usecase *userUsecase) GetAllData(ctx context.Context) (data []users.Core, err error) {
	data, err = usecase.userData.GetAllData(ctx)
	return data, err
}

func (usecase *userUsecase) GetData(ctx context.Context, id int) (data users.Core, err error) {
	data, err = usecase.userData.GetData(ctx, id)
	return data, err
}

func (usecase *userUsecase) DeleteData(ctx context.Context, id int) (row int, err error) {
	row, err = usecase.userData.DeleteData(ctx, id)
	return row, err
}

func (usecase *userUsecase) UpdateData(ctx context.Context, id int, insert users.Core) (row int, err error) {
	row, err = usecase.userData.UpdateData(ctx, id, insert)
	return row, err
}
