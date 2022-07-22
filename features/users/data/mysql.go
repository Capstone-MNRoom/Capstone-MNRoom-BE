package data

import (
	"be9/mnroom/features/users"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	// "google.golang.org/protobuf/internal/errors"

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
	if tx.Error.(*mysql.MySQLError).Number == 1062 {
		return 0, errors.New("email already exist")
	}

	// if tx.Error != nil {
	// 	return 0, tx.Error
	// }
	// if tx.RowsAffected != 1 {
	// 	return 0, fmt.Errorf("failed to insert data")
	// }
	return int(tx.RowsAffected), nil
}

func (repo *mysqlUserRepository) GetAllData(limit int, offset int) (data []users.Core, err error) {
	var getAllData []User
	tx := repo.db.Limit(limit).Offset(offset).Find(&getAllData)
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
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to deleted data")
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlUserRepository) UpdateData(id int, insert users.Core) (row int, err error) {
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(User{Image: insert.Image, Username: insert.Username, Email: insert.Email, Password: insert.Password, Phone: insert.Phone, Address: insert.Address})
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected != 1 {
		return 0, fmt.Errorf("failed to updated data")
	}
	return int(tx.RowsAffected), nil
}
