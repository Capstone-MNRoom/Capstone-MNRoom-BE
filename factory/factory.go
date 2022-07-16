package factory

import (
	_userBusiness "be9/mnroom/features/users/business"
	_userData "be9/mnroom/features/users/data"
	_userPresentation "be9/mnroom/features/users/presentation"

	_authBusiness "be9/mnroom/features/login/business"
	_authData "be9/mnroom/features/login/data"
	_authPresentation "be9/mnroom/features/login/presentation"

	_categoryBusiness "be9/mnroom/features/categorys/business"
	_categoryData "be9/mnroom/features/categorys/data"
	_categoryPresentation "be9/mnroom/features/categorys/presentation"

	_facilityBusiness "be9/mnroom/features/facilitys/business"
	_facilityData "be9/mnroom/features/facilitys/data"
	_facilityPresentation "be9/mnroom/features/facilitys/presentation"

	_roomBusiness "be9/mnroom/features/rooms/business"
	_roomData "be9/mnroom/features/rooms/data"
	_roomPresentation "be9/mnroom/features/rooms/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	// Login
	AuthPresenter *_authPresentation.AuthHandler
	// Users
	UserPresenter *_userPresentation.UserHandler
	// Categorys
	CategoryPresenter *_categoryPresentation.CategoryHandler
	// Facilitys
	FacilityPresenter *_facilityPresentation.FacilityHandler
	// Rooms
	RoomPresenter *_roomPresentation.RoomHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	authData := _authData.NewAuthRepository(dbConn)
	authBusiness := _authBusiness.NewAuthBusiness(authData)
	authPresentation := _authPresentation.NewAuthHandler(authBusiness)

	categoryData := _categoryData.NewCategoryRepository(dbConn)
	categoryBusiness := _categoryBusiness.NewCategoryBusiness(categoryData)
	categoryPresentation := _categoryPresentation.NewCategoryHandler(categoryBusiness)

	facilityData := _facilityData.NewFacilityRepository(dbConn)
	facilityBusiness := _facilityBusiness.NewFacilityBusiness(facilityData)
	facilityPresentation := _facilityPresentation.NewFacilityHandler(facilityBusiness)

	roomData := _roomData.NewRoomRepository(dbConn)
	roomBusiness := _roomBusiness.NewRoomBusiness(roomData)
	roomPresentation := _roomPresentation.NewRoomHandler(roomBusiness)

	return Presenter{
		UserPresenter:     userPresentation,
		AuthPresenter:     authPresentation,
		CategoryPresenter: categoryPresentation,
		FacilityPresenter: facilityPresentation,
		RoomPresenter:     roomPresentation,
	}
}
