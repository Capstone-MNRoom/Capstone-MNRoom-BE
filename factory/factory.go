package factory

import (
	_userBusiness "be9/mnroom/features/users/business"
	_userData "be9/mnroom/features/users/data"
	_userPresentation "be9/mnroom/features/users/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	// Users
	UserPresenter *_userPresentation.UserHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	return Presenter{
		UserPresenter: userPresentation,
	}
}
