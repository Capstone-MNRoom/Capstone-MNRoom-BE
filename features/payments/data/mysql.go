package data

import (
	"be9/mnroom/features/payments"
	_rents "be9/mnroom/features/rents/data"
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
	var payment Payments
	tx := repo.db.WithContext(ctx).Where("order_id = ?", idOrder).First(&payment)
	if tx.Error != nil {
		return 0, tx.Error
	}

	payment.TransactionStatus = status
	if err := repo.db.WithContext(ctx).Save(&payment).Error; err != nil {
		return 0, err
	}

	// Synchronize Rents status upon payment completion
	if status == "success" || status == "settlement" {
		repo.db.WithContext(ctx).Model(&_rents.Rents{}).Where("id = ?", payment.RentsID).Update("status", "Success")
	} else if status == "failed" {
		repo.db.WithContext(ctx).Model(&_rents.Rents{}).Where("id = ?", payment.RentsID).Update("status", "Failed")
	}

	return 1, nil
}
