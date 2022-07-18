package data

import (
	"be9/mnroom/features/roomfacilitys"

	"gorm.io/gorm"
)

type mysqlRoomFacilityRepository struct {
	db *gorm.DB
}

func NewRoomFacilityRepository(conn *gorm.DB) roomfacilitys.Data {
	return &mysqlRoomFacilityRepository{
		db: conn,
	}
}
