package data

import (
	"be9/mnroom/features/login"
)

type User struct {
	ID       uint   `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (data *User) toCore() login.Core {
	return login.Core{
		ID:       int(data.ID),
		Email:    data.Email,
		Password: data.Password,
	}
}

func fromCore(core login.Core) User {
	return User{
		Email:    core.Email,
		Password: core.Password,
	}
}
