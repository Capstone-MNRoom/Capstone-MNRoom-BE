package data

import (
	"be9/mnroom/features/rents"
	_rooms "be9/mnroom/features/rooms/data"
	"context"

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

func (repo *mysqlRentRepository) GetData(ctx context.Context, id int) (data int, err error) {
	var getData _rooms.Rooms
	tx := repo.db.WithContext(ctx).First(&getData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return getData.RentalPrice, nil
}

func (repo *mysqlRentRepository) GetDataRentToken(ctx context.Context, idToken int, idRoom int) (row int, err error) {
	var getData Rents
	tx := repo.db.WithContext(ctx).Where("user_id = ? AND rooms_id = ?", idToken, idRoom).Preload("User").Preload("Rooms").First(&getData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRentRepository) GetDataRentUser(ctx context.Context, id int, start string, end string) (row int, err error) {
	var getData Rents
	tx := repo.db.WithContext(ctx).Where("rooms_id = ? AND date_start AND date_end BETWEEN ? AND ?", id, start, end).Preload("User").Preload("Rooms").Find(&getData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRentRepository) InsertData(ctx context.Context, insert rents.Core) (row int, err error) {
	insertData := fromCore(insert)
	tx := repo.db.WithContext(ctx).Create(&insertData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(insertData.ID), nil
}

func (repo *mysqlRentRepository) GetDataRent(ctx context.Context, id int) (data []rents.Core, err error) {
	var getData []Rents
	tx := repo.db.WithContext(ctx).Where("rooms_id = ?", id).Preload("User").Preload("Rooms").Find(&getData)
	if tx.Error != nil {
		return []rents.Core{}, tx.Error
	}
	return toCoreList(getData), nil
}

func (repo *mysqlRentRepository) InsertDataPayment(ctx context.Context, insert rents.CorePayments) (data rents.CorePayments, err error) {
	insertData := fromCorePayment(insert)
	tx := repo.db.WithContext(ctx).Create(&insertData)
	if tx.Error != nil {
		return rents.CorePayments{}, tx.Error
	}
	return insertData.toCorePayment(), nil
}

func (repo *mysqlRentRepository) GetDataUser(ctx context.Context, idToken int) (data rents.CoreUser, err error) {
	var getData User
	tx := repo.db.WithContext(ctx).First(&getData, idToken)
	if tx.Error != nil {
		return rents.CoreUser{}, tx.Error
	}
	return rents.CoreUser(getData.toCoreUser()), nil
}

func (repo *mysqlRentRepository) GetDataRentUserHistory(ctx context.Context, idToken int) (data []rents.Core, err error) {
	var getDataUser []Rents
	tx := repo.db.WithContext(ctx).Where("user_id = ?", idToken).Preload("User").Preload("Rooms").Find(&getDataUser)
	if tx.Error != nil {
		return []rents.Core{}, tx.Error
	}
	return toCoreList(getDataUser), nil
}
