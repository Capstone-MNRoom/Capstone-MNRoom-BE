package data

import (
	"be9/mnroom/features/payments"

	"gorm.io/gorm"
)

type mysqlPaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(conn *gorm.DB) payments.Data {
	return &mysqlPaymentRepository{
		db: conn,
	}
}

func (repo *mysqlPaymentRepository) GetAllData(idToken int) (data []payments.Core, err error) {
	var getData []Payments
	tx := repo.db.Where("user_id = ?", idToken).Preload("User").Preload("Rents").Find(&getData)
	if tx.Error != nil {
		return []payments.Core{}, nil
	}
	return toCoreList(getData), nil
}
