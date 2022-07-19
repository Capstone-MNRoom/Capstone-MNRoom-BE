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

func (rnt *rentUseCase) GetDataRentToken(idToken int) (row int, err error) {
	row, err = rnt.rentData.GetDataRentToken(idToken)
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
