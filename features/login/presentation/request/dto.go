package request

import (
	"be9/mnroom/features/login"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(req User) login.Core {
	return login.Core{
		Email:    req.Email,
		Password: req.Password,
	}
}
