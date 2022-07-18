package data

import (
	"be9/mnroom/features/rents"

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

func (repo *mysqlRentRepository) InsertData(insert rents.Core) (row int, err error) {
	insertRent := fromCore(insert)
	tx := repo.db.Create(&insertRent)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
