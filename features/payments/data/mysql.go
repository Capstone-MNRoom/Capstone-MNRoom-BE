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
