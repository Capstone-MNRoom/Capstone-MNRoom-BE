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

func (uc *rentUseCase) InsertData(insert rents.Core) (row int, err error) {
	row, err = uc.rentData.InsertData(insert)
	return row, err
}
