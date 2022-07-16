package data

import (
	"be9/mnroom/features/rooms"
	"fmt"

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
	insertData := fromCore(insert)
	tx := repo.db.Create(&insertData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to insert data")
	}
	return int(tx.RowsAffected), nil
}

// func (repo *mysqlRoomRepository) GetAllData(limit int, offset int) (data []rooms.Core, err error) {
// 	var getAllData []Rooms
// 	tx := repo.db.Limit(limit).Offset(offset).Preload("User").Preload("Categorys").Preload("Facilitys").Find(&getAllData)
// 	if tx.Error != nil {
// 		return []rooms.Core{}, tx.Error
// 	}
// 	return toCoreList(getAllData), nil
// }

// func (repo *mysqlRoomRepository) GetData(id int) (data rooms.Core, err error) {
// 	var getData Rooms
// 	tx := repo.db.Preload("User").Preload("Categorys").Preload("Facilitys").First(&getData, id)
// 	if tx.Error != nil {
// 		return rooms.Core{}, tx.Error
// 	}
// 	return getData.toCore(), nil
// }

// func (repo *mysqlRoomRepository) DeleteData(id int) (row int, err error) {
// 	var deleteData Rooms
// 	tx := repo.db.Unscoped().Delete(&deleteData, id)
// 	if tx.Error != nil {
// 		return 0, tx.Error
// 	}
// 	if tx.RowsAffected != 1 {
// 		return 0, fmt.Errorf("failed to deleted data")
// 	}
// 	return int(tx.RowsAffected), nil
// }
