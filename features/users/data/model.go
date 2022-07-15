package data

import (
	"be9/mnroom/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Image    string `json:"image" form:"image"`
	Username string `json:"username" form:"username"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `gorm:"unique" json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}

func toCoreList(data []User) []users.Core {
	result := []users.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *User) toCore() users.Core {
	return users.Core{
		ID:        int(data.ID),
		Image:     data.Image,
		Username:  data.Username,
		Email:     data.Email,
		Password:  data.Password,
		Phone:     data.Phone,
		Address:   data.Address,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func fromCore(core users.Core) User {
	return User{
		Image:    core.Image,
		Username: core.Username,
		Email:    core.Email,
		Password: core.Password,
		Phone:    core.Phone,
		Address:  core.Address,
	}
}
