package business

import (
	"be9/mnroom/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockUserData struct{}

// GetAllData implements users.Data
func (mockUserData) GetAllData() (data []users.Core, err error) {
	panic("unimplemented")
}

// GetData implements users.Data
func (mockUserData) GetData(id int) (data users.Core, err error) {
	panic("unimplemented")
}

// InsertData implements users.Data
func (mockUserData) InsertData(insert users.Core) (row int, err error) {
	panic("unimplemented")
}

// UpdateData implements users.Data
func (mockUserData) UpdateData(id int, insert users.Core) (row int, err error) {
	panic("unimplemented")
}

// InsertData(insert Core) (row int, err error)
// func (mock mockUserData) InsertData(data users.Core) (row int, err error) {
// 	return 1, nil
// }

// func (mock mockUserData) UpdateData(id int, data users.Core) (row int, err error) {
// 	return 1, nil
// }

// func (mock mockUserData) GetData(id int) (data users.Core, err error) {
// 	return users.Core{
// 		ID:       1,
// 		Image:    "https://storage.googleapis.com/profile/profile_default.png",
// 		Username: "Mulya Nurdin",
// 		Email:    "mulya@mail.com",
// 		Phone:    "081234567890",
// 		Address:  "Bandung",
// 	}, nil
// }
func (mock mockUserData) DeleteData(id int) (row int, err error) {
	return 1, nil
}

// Mock Failed
type mockUserDataFailed struct{}

// GetAllData implements users.Data
func (mockUserDataFailed) GetAllData() (data []users.Core, err error) {
	panic("unimplemented")
}

// GetData implements users.Data
func (mockUserDataFailed) GetData(id int) (data users.Core, err error) {
	panic("unimplemented")
}

// InsertData implements users.Data
func (mockUserDataFailed) InsertData(insert users.Core) (row int, err error) {
	panic("unimplemented")
}

// UpdateData implements users.Data
func (mockUserDataFailed) UpdateData(id int, insert users.Core) (row int, err error) {
	panic("unimplemented")
}

func (mockUserDataFailed) DeleteData(id int) (row int, err error) {
	return 0, nil
}


func TestInsertData(t *testing.T) {
	t.Run("Test insert data user", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		newUser := users.Core{
			Username: "Mulya Nurdin",
			Email:    "mei@mail.com",
			Password: "qwerty1234",
			Phone:    "031123456",
			Address:  "Bandung",
		}
		result, err := userBusiness.InsertData(newUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
}

func TestGetData(t *testing.T) {
	t.Run("Test get data user", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.GetData(1)
		assert.Nil(t, err)
		assert.Equal(t, "mulya", result.Username)
	})
}

func TestDeleteData(t *testing.T) {
	t.Run("Test Delete User Succes", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.DeleteData(0)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
	})
	t.Run("Test Delete User Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.DeleteData(0)
		assert.Nil(t, err)
		assert.Equal(t, 0, result)
	})
}
