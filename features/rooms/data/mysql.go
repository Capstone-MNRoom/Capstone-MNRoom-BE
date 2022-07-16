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
