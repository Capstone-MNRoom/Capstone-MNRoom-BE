package migration

import (
	_mCategorys "be9/mnroom/features/categorys/data"
	_mFacilitys "be9/mnroom/features/facilitys/data"
	_mRents "be9/mnroom/features/rents/data"
	_mRooms "be9/mnroom/features/rooms/data"
	_mUsers "be9/mnroom/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mCategorys.Categorys{})
	db.AutoMigrate(&_mFacilitys.Facilitys{})
	db.AutoMigrate(&_mRooms.Rooms{})
	db.AutoMigrate(&_mRents.Rents{})

}
