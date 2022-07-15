package request

import "be9/mnroom/features/login"

type User struct {
	Email    string `json:"email" validate:"required,email" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
}

func ToCore(req User) login.Core {
	return login.Core{
		Email:    req.Email,
		Password: req.Password,
	}
}
