package data

import (
	"be9/mnroom/features/login"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(conn *gorm.DB) login.Data {
	return &mysqlAuthRepository{
		db: conn,
	}
}

func (repo *mysqlAuthRepository) Auth(data login.Core) (dataAuth login.Core, err error) {
	var authUser User
	tx := repo.db.First(&authUser, "email = ?", data.Email)
	if tx.Error != nil {
		return login.Core{}, tx.Error
	}
	err = bcrypt.CompareHashAndPassword([]byte(authUser.toCore().Password), []byte(data.Password))
	if err != nil {
		return login.Core{}, err
	}
	return authUser.toCore(), nil
}
