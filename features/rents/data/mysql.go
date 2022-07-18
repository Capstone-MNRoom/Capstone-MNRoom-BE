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
