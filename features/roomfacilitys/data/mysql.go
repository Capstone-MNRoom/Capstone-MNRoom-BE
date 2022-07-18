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

func (repo *mysqlRoomFacilityRepository) GetData(id int) (data []roomfacilitys.Core, err error) {
	var getData []RoomFacilitys
	tx := repo.db.Where("rooms_id = ?", id).Preload("User").Preload("Rooms").Preload("Facilitys").Find(&getData)
	if tx.Error != nil {
		return []roomfacilitys.Core{}, err
	}
	return toCoreList(getData), nil
}
