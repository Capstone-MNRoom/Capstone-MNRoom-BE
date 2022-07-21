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
		return []roomfacilitys.Core{}, tx.Error
	}
	return toCoreList(getData), nil
}

func (repo *mysqlRoomFacilityRepository) GetDataRow(id int) (row int, err error) {
	var getData RoomFacilitys
	tx := repo.db.Where("rooms_id = ?", id).Preload("User").Preload("Rooms").Preload("Facilitys").First(&getData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
