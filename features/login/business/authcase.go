package business

import (
	"be9/mnroom/features/login"
	"context"
)

type authCase struct {
	authData login.Data
}

func NewAuthBusiness(athData login.Data) login.Business {
	return &authCase{
		authData: athData,
	}
}

func (ac *authCase) Auth(ctx context.Context, data login.Core) (dataAuth login.Core, err error) {
	dataAuth, err = ac.authData.Auth(ctx, data)
	return dataAuth, err
}
