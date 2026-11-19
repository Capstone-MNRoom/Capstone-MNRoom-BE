package data

import (
	"be9/mnroom/features/users"
	"context"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) InsertData(ctx context.Context, insert users.Core) (row int, err error) {
	insertData := fromCore(insert)
	tx := repo.db.WithContext(ctx).Create(&insertData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlUserRepository) GetAllData(ctx context.Context) (data []users.Core, err error) {
	var getAllData []User
	tx := repo.db.WithContext(ctx).Find(&getAllData)
	if tx.Error != nil {
		return []users.Core{}, tx.Error
	}
	return toCoreList(getAllData), nil
}

func (repo *mysqlUserRepository) GetData(ctx context.Context, id int) (data users.Core, err error) {
	var getData User
	tx := repo.db.WithContext(ctx).First(&getData, id)
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}
	return getData.toCore(), nil
}

func (repo *mysqlUserRepository) DeleteData(ctx context.Context, id int) (row int, err error) {
	var getData User
	tx := repo.db.WithContext(ctx).Unscoped().Delete(&getData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlUserRepository) UpdateData(ctx context.Context, id int, insert users.Core) (row int, err error) {
	tx := repo.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Updates(User{Image: insert.Image, Username: insert.Username, Password: insert.Password, Address: insert.Address})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
