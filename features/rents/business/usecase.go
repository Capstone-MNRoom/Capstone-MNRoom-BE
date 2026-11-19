package business

import (
	"be9/mnroom/features/rents"
	"context"
)

type rentUseCase struct {
	rentData rents.Data
}

func NewRentBusiness(rent rents.Data) rents.Business {
	return &rentUseCase{
		rentData: rent,
	}
}

func (rnt *rentUseCase) GetData(ctx context.Context, id int) (data int, err error) {
	data, err = rnt.rentData.GetData(ctx, id)
	return data, err
}

func (rnt *rentUseCase) InsertData(ctx context.Context, insert rents.Core) (row int, err error) {
	row, err = rnt.rentData.InsertData(ctx, insert)
	return row, err
}

func (rnt *rentUseCase) GetDataRentToken(ctx context.Context, idToken int, idRoom int) (row int, err error) {
	row, err = rnt.rentData.GetDataRentToken(ctx, idToken, idRoom)
	return row, err
}

func (rnt *rentUseCase) GetDataRentUser(ctx context.Context, id int, start string, end string) (row int, err error) {
	row, err = rnt.rentData.GetDataRentUser(ctx, id, start, end)
	return row, err
}

func (rnt *rentUseCase) GetDataRent(ctx context.Context, id int) (data []rents.Core, err error) {
	data, err = rnt.rentData.GetDataRent(ctx, id)
	return data, err
}

func (rnt *rentUseCase) InsertDataPayment(ctx context.Context, insert rents.CorePayments) (data rents.CorePayments, err error) {
	data, err = rnt.rentData.InsertDataPayment(ctx, insert)
	return data, err
}

func (rnt *rentUseCase) GetDataUser(ctx context.Context, idToken int) (data rents.CoreUser, err error) {
	data, err = rnt.rentData.GetDataUser(ctx, idToken)
	return data, err
}

func (rnt *rentUseCase) GetDataRentUserHistory(ctx context.Context, idToken int) (data []rents.Core, err error) {
	data, err = rnt.rentData.GetDataRentUserHistory(ctx, idToken)
	return data, err
}
