package business

import "be9/mnroom/features/login"

type authCase struct {
	authData login.Data
}

func NewAuthBusiness(athData login.Data) login.Business {
	return &authCase{
		authData: athData,
	}
}

func (ac *authCase) Auth(data login.Core) (dataAuth login.Core, err error) {
	dataAuth, err = ac.authData.Auth(data)
	return dataAuth, err
}
