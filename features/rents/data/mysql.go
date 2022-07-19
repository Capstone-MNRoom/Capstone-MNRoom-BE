package data

import (
	"be9/mnroom/features/rents"
	_rooms "be9/mnroom/features/rooms/data"

	"gorm.io/gorm"
)

type mysqlRentRepository struct {
	db *gorm.DB
}

func NewRentRepository(conn *gorm.DB) rents.Data {
	return &mysqlRentRepository{
		db: conn,
	}
}

func (repo *mysqlRentRepository) GetData(id int) (data int, err error) {
	var getData _rooms.Rooms
	tx := repo.db.First(&getData, id)
	if tx.Error != nil {
		return 0, err
	}
	return getData.RentalPrice, nil
}

func (repo *mysqlRentRepository) InsertData(insert rents.Core) (row int, err error) {
	insertData := fromCore(insert)
	tx := repo.db.Create(&insertData)
	if tx.Error != nil {
		return 0, err
	}
	return int(tx.RowsAffected), nil
}
