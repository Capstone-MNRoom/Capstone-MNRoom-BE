package business

import (
	"be9/mnroom/features/users"
	"fmt"
	// "testing"

	// "github.com/stretchr/testify/assert"
	// "be9/mnroom/features/users"
	// "fmt"
	// "testing"
	// "github.com/stretchr/testify/assert"
)

type mockUserData struct{}

// InsertData(insert Core) (row int, err error)
func (mock mockUserData) InsertData(data users.Core) (row int, err error) {
	return 1, nil
}

func (mock mockUserData) UpdateData(id int, data users.Core) (row int, err error) {
	return 1, nil
}

func (mock mockUserData) GetData(id int) (data users.Core, err error) {
	return users.Core{
		ID:       1,
		Image:    "https://storage.googleapis.com/profile/profile_default.png",
		Username: "Mulya Nurdin",
		Email:    "mulya@mail.com",
		Phone:    "081234567890",
		Address:  "Bandung",
	}, nil
}
func (mock mockUserData) DeleteData(id int) (row int, err error) {
	return 1, nil
}

type mockUserDataFailed struct {
}

func (mock mockUserDataFailed) InsertData(data users.Core) (row int, err error) {
	return 0, fmt.Errorf("failed insert data")
}

// func (mock mockUserDataFailed) AuthUser(email string, password string) (username string, token string, err error) {
// 	return "", "", fmt.Errorf("failed")
// }
func (mock mockUserDataFailed) UpdateData(id int, data users.Core) (row int, err error) {
	return 0, fmt.Errorf("failed insert data")
}
func (mock mockUserDataFailed) GetData(id int) (data users.Core, err error) {
	return users.Core{}, fmt.Errorf("error")
}
func (mock mockUserDataFailed) GetAllData(id int) (data users.Core, err error) {
	return users.Core{}, fmt.Errorf("error")
}
func (mock mockUserDataFailed) DeleteData(id int) (row int, err error) {
	return 0, nil
}

// type Business interface {
// 	InsertData(insert Core) (row int, err error)
// 	GetAllData(limit int, offset int) (data []Core, err error)
// 	GetData(id int) (data Core, err error)
// 	DeleteData(id int) (row int, err error)
// 	UpdateData(id int, insert Core) (row int, err error)
// }

// func TestInsertData(t *testing.T) {
// 	t.Run("Test insert data user", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserData{})
// 		newUser := users.Core{
// 			Username: "Mulya Nurdin",
// 			Email:    "mei@mail.com",
// 			Password: "qwerty1234",
// 			Phone:    "031123456",
// 			Address:  "Bandung",
// 		}
// 		result, err := userBusiness.InsertData(newUser)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, result)
// 	})
// }

// func TestGetData(t *testing.T) {
// 	t.Run("Test get data user", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserData)
// 		result, err := userBusiness.GetData(0)
// 		assert.Nil(t, err)
// 		assert.Equal(t, "mulya", result.Username)
// 	})
// }
