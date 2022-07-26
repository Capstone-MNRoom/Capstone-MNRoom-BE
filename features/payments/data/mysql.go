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

func (repo *mysqlPaymentRepository) UpdateData(idOrder string, status string) (row int, err error) {
	tx := repo.db.Model(&Payments{}).Where("order_id = ?", idOrder).Update("transaction_status", status)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
