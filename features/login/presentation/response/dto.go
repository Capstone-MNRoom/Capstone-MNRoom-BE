package response

import "be9/mnroom/features/login"

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func FromCore(data login.Core) User {
	return User{
		ID:    data.ID,
		Email: data.Email,
	}
}
