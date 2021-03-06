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

	_roomfacilitysBusiness "be9/mnroom/features/roomfacilitys/business"
	_roomfacilitysData "be9/mnroom/features/roomfacilitys/data"
	_roomfacilitysPresentation "be9/mnroom/features/roomfacilitys/presentation"

	_rentBusiness "be9/mnroom/features/rents/business"
	_rentData "be9/mnroom/features/rents/data"
	_rentPresentation "be9/mnroom/features/rents/presentation"

	_feedbackBusiness "be9/mnroom/features/feedback/business"
	_feedbackData "be9/mnroom/features/feedback/data"
	_feedbackPresentation "be9/mnroom/features/feedback/presentation"

	_paymentBusiness "be9/mnroom/features/payments/business"
	_paymentData "be9/mnroom/features/payments/data"
	_paymentPresentation "be9/mnroom/features/payments/presentation"

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
	// Room Facilitys
	RoomFacilitysPresenter *_roomfacilitysPresentation.RoomFacilityHandler
	// Rents
	RentPresenter *_rentPresentation.RentHandler
	// Feedback
	FeedbackPresenter *_feedbackPresentation.FeedbackHandler
	// Payments
	PaymentPresenter *_paymentPresentation.PaymentHandler
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

	roomfacilityData := _roomfacilitysData.NewRoomFacilityRepository(dbConn)
	roomfacilityBusiness := _roomfacilitysBusiness.NewRoomFacilityBusiness(roomfacilityData)
	roomfacilityPresentation := _roomfacilitysPresentation.NewRoomFacilitysHandler(roomfacilityBusiness)

	rentData := _rentData.NewRentRepository(dbConn)
	rentBusiness := _rentBusiness.NewRentBusiness(rentData)
	rentPresentation := _rentPresentation.NewEventHandler(rentBusiness)

	feedbackData := _feedbackData.NewFeedbackRepository(dbConn)
	feedbackBusiness := _feedbackBusiness.NewFeedbackBusiness(feedbackData)
	feedbackPresentation := _feedbackPresentation.NewFeedbackHandler(feedbackBusiness)

	paymentData := _paymentData.NewPaymentRepository(dbConn)
	paymentBusiness := _paymentBusiness.NewPaymentBusiness(paymentData)
	paymentPresentation := _paymentPresentation.NewPaymentHandler(paymentBusiness)

	return Presenter{
		UserPresenter:          userPresentation,
		AuthPresenter:          authPresentation,
		CategoryPresenter:      categoryPresentation,
		FacilityPresenter:      facilityPresentation,
		RoomPresenter:          roomPresentation,
		RoomFacilitysPresenter: roomfacilityPresentation,
		RentPresenter:          rentPresentation,
		FeedbackPresenter:      feedbackPresentation,
		PaymentPresenter:       paymentPresentation,
	}
}
