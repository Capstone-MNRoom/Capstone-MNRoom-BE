package response

import (
	"be9/mnroom/features/users"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Image     string    `json:"image"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

func FromCoreList(data []users.Core) []User {
	result := []User{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data users.Core) User {
	return User{
		ID:        data.ID,
		Image:     data.Image,
		Username:  data.Username,
		Email:     data.Email,
		Phone:     data.Phone,
		Address:   data.Address,
		CreatedAt: data.CreatedAt,
	}
}
