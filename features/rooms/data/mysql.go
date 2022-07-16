package data

import (
	"be9/mnroom/features/rooms"

	"gorm.io/gorm"
)

type mysqlRoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(conn *gorm.DB) rooms.Data {
	return &mysqlRoomRepository{
		db: conn,
	}
}

func (repo *mysqlRoomRepository) InsertData(insert rooms.Core) (row int, err error) {
	insertRoom := fromCore(insert)
	tx := repo.db.Create(&insertRoom)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRoomRepository) GetDataAll() (data []rooms.Core, err error) {
	var getDataAll []Rooms
	tx := repo.db.Preload("User").Preload("Categorys").Find(&getDataAll)
	if tx.Error != nil {
		return []rooms.Core{}, tx.Error
	}
	return toCoreList(getDataAll), nil
}

func (repo *mysqlRoomRepository) GetData(id int) (data rooms.Core, err error) {
	var getData Rooms
	tx := repo.db.Preload("User").Preload("Categorys").First(&getData, id)
	if tx.Error != nil {
		return rooms.Core{}, tx.Error
	}
	return getData.toCore(), nil
}
