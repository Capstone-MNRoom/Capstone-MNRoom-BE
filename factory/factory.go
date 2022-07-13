package factory

import (
	"gorm.io/gorm"
)

type Presenter struct {
}

func InitFactory(dbConn *gorm.DB) Presenter {

	return Presenter{}
}
