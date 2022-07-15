package factory

import (
	_userBusiness "be9/mnroom/features/users/business"
	_userData "be9/mnroom/features/users/data"
	_userPresentation "be9/mnroom/features/users/presentation"

	_authBusiness "be9/mnroom/features/login/business"
	_authData "be9/mnroom/features/login/data"
	_authPresentation "be9/mnroom/features/login/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	// Login
	AuthPresenter *_authPresentation.AuthHandler
	// Users
	UserPresenter *_userPresentation.UserHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	authData := _authData.NewAuthRepository(dbConn)
	authBusiness := _authBusiness.NewAuthBusiness(authData)
	authPresentation := _authPresentation.NewAuthHandler(authBusiness)

	return Presenter{
		UserPresenter: userPresentation,
		AuthPresenter: authPresentation,
	}
}
