package request

import "be9/mnroom/features/users"

type User struct {
	Image    string `json:"image" form:"image"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}

func ToCore(req User) users.Core {
	return users.Core{
		Image:    req.Image,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		Address:  req.Address,
	}
}
