package business

import "be9/mnroom/features/rents"

type rentUseCase struct {
	rentData rents.Data
}

func NewRentBusiness(rent rents.Data) rents.Business {
	return &rentUseCase{
		rentData: rent,
	}
}

func (rnt *rentUseCase) GetData(id int) (data int, err error) {
	data, err = rnt.rentData.GetData(id)
	return data, err
}

func (rnt *rentUseCase) InsertData(insert rents.Core) (row int, err error) {
	row, err = rnt.rentData.InsertData(insert)
	return row, err
}

func (rnt *rentUseCase) GetDataRentToken(idToken int, idRoom int) (row int, err error) {
	row, err = rnt.rentData.GetDataRentToken(idToken, idRoom)
	return row, err
}

func (rnt *rentUseCase) GetDataRentUser(id int, start string, end string) (row int, err error) {
	row, err = rnt.rentData.GetDataRentUser(id, start, end)
	return row, err
}

func (rnt *rentUseCase) GetDataRent(id int) (data []rents.Core, err error) {
	data, err = rnt.rentData.GetDataRent(id)
	return data, err
}

func (rnt *rentUseCase) InsertDataPayment(insert rents.CorePayments) (data rents.CorePayments, err error) {
	data, err = rnt.rentData.InsertDataPayment(insert)
	return data, err
}

func (rnt *rentUseCase) GetDataUser(idToken int) (data rents.CoreUser, err error) {
	data, err = rnt.rentData.GetDataUser(idToken)
	return data, err
}

func (rnt *rentUseCase) GetDataRentUserHistory(idToken int) (data []rents.Core, err error) {
	data, err = rnt.rentData.GetDataRentUserHistory(idToken)
	return data, err
}
