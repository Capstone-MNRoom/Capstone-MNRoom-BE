package data

import (
	"be9/mnroom/features/users"

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

func (repo *mysqlUserRepository) InsertData(insert users.Core) (row int, err error) {
	insertData := fromCore(insert)
	tx := repo.db.Create(&insertData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlUserRepository) GetAllData() (data []users.Core, err error) {
	var getAllData []User
	tx := repo.db.Find(&getAllData)
	if tx.Error != nil {
		return []users.Core{}, tx.Error
	}
	return toCoreList(getAllData), nil
}

func (repo *mysqlUserRepository) GetData(id int) (data users.Core, err error) {
	var getData User
	tx := repo.db.First(&getData, id)
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}
	return getData.toCore(), nil
}

func (repo *mysqlUserRepository) DeleteData(id int) (row int, err error) {
	var getData User
	tx := repo.db.Unscoped().Delete(&getData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlUserRepository) UpdateData(id int, insert users.Core) (row int, err error) {
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(User{Image: insert.Image, Username: insert.Username, Password: insert.Password, Address: insert.Address})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
