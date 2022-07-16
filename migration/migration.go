package migration

import (
	_mCategorys "be9/mnroom/features/categorys/data"
	_mUsers "be9/mnroom/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mCategorys.Categorys{})
}
