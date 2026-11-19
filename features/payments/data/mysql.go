package data

import (
	"be9/mnroom/features/payments"
	"context"

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

func (repo *mysqlPaymentRepository) GetAllData(ctx context.Context, idToken int) (data []payments.Core, err error) {
	var getData []Payments
	tx := repo.db.WithContext(ctx).Where("user_id = ?", idToken).Preload("User").Preload("Rents").Find(&getData)
	if tx.Error != nil {
		return []payments.Core{}, nil
	}
	return toCoreList(getData), nil
}

func (repo *mysqlPaymentRepository) UpdateData(ctx context.Context, idOrder string, status string) (row int, err error) {
	tx := repo.db.WithContext(ctx).Model(&Payments{}).Where("order_id = ?", idOrder).Update("transaction_status", status)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
