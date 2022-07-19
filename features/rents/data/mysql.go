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
		return 0, tx.Error
	}
	return getData.RentalPrice, nil
}

func (repo *mysqlRentRepository) GetDataRentToken(idToken int) (row int, err error) {
	var getData Rents
	tx := repo.db.Where("user_id = ?", idToken).Preload("User").Preload("Rooms").First(&getData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRentRepository) GetDataRentUser(id int, start string, end string) (row int, err error) {
	var getData Rents
	tx := repo.db.Where("rooms_id = ? AND date_start = ? AND date_end = ?", id, start, end).Preload("User").Preload("Rooms").First(&getData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRentRepository) InsertData(insert rents.Core) (row int, err error) {
	insertData := fromCore(insert)
	tx := repo.db.Create(&insertData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRentRepository) GetDataRent(id int) (data []rents.Core, err error) {
	var getData []Rents
	tx := repo.db.Where("rooms_id = ?", id).Preload("User").Preload("Rooms").Find(&getData)
	if tx.Error != nil {
		return []rents.Core{}, tx.Error
	}
	return toCoreList(getData), nil
}
