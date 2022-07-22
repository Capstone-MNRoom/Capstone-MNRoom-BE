package business

// import (
// 	"be9/mnroom/features/users"
// 	"fmt"
// 	"testing"
// 	// "time"

// 	// "github.com/stretchr/testify/assert"
// )

// type mockUserData struct{}

// func (mock mockUserData) PostUser(data users.Core) (row int, err error) {
// 	return 1, nil
// }

// func (mock mockUserData) AuthUser(email string, password string) (userName string, token string, e error) {
// 	return "andri", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", nil
// }

// func (mock mockUserData) PutDataUser(id int, data users.Core) (row int, err error) {
// 	return 1, nil
// }

// func (mock mockUserData) GetUser(id int) (data users.Core, err error) {
// 	return users.Core{
// 		ID: 1, Username: "Mulya Nurdin", Email: "mulya@mail.com", PhoneNumber
// 		//Username: "Andri G", UserName: "Andri", Email: "andri@gmail.com", PhoneNumber: "081234", Address: "Jln. Urip", ImageURL: "https://storage.com/profile.png", CreatedAt: time.Time{},
// 	}, nil
// }

// func (mockUserData) DeleteUser(id int) (row int, err error) {
// 	return 1, nil
// }

// type mockUserDataFailed struct{}

// func (mock mockUserDataFailed) PostUser(data users.Core) (row int, err error) {
// 	return 0, fmt.Errorf("Failed to select data")
// }

// func (mock mockUserDataFailed) AuthUser(email string, password string) (userName string, token string, e error) {
// 	return "", "", fmt.Errorf("email or password incorrect")
// }

// func (mock mockUserDataFailed) PutDataUser(id int, data users.Core) (row int, err error) {
// 	return 0, fmt.Errorf("failed to update user")
// }

// func (mock mockUserDataFailed) GetUser(id int) (data users.Core, err error) {
// 	return users.Core{}, fmt.Errorf("failed to get data")
// }

// func (mockUserDataFailed) DeleteUser(id int) (row int, err error) {
// 	return 0, nil
// }

// func TestInsertUser(t *testing.T) {
// 	t.Run("Test Insert User Success", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserData{})
// 		newUser := users.Core{
// 			Us:    "Andri",
// 			Email:       "andri@gmail.com",
// 			Password:    "qwerty",
// 			PhoneNumber: "081234",
// 			Address:     "Jalan Urip",
// 			ImageURL:    "profile.png",
// 		}
// 		result, err := userBusiness.InsertUser(newUser)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, result)
// 	})

// 	t.Run("Test Insert Data Failed", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		newUser := users.Core{
// 			FullName:    "Andri G",
// 			UserName:    "Andri",
// 			Email:       "andri@gmail.com",
// 			Password:    "qwerty",
// 			PhoneNumber: "081234",
// 			Address:     "Jalan Urip",
// 			ImageURL:    "profile.png",
// 		}
// 		result, err := userBusiness.InsertUser(newUser)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, 0, result)
// 	})
// 	t.Run("Test Insert Data Failed When Full Name Empty", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		newUser := users.Core{
// 			UserName:    "Andri",
// 			Email:       "andri@gmail.com",
// 			Password:    "qwerty",
// 			PhoneNumber: "081234",
// 			Address:     "Jalan Urip",
// 			ImageURL:    "profile.png",
// 		}
// 		result, err := userBusiness.InsertUser(newUser)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, -1, result)
// 	})
// }

// func TestLoginUser(t *testing.T) {
// 	t.Run("Test Login User Success", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserData{})
// 		var loginEmail = "andri@gmail.com"
// 		var loginPass = "qwerty"

// 		resUserName, resToken, err := userBusiness.LoginUser(loginEmail, loginPass)
// 		assert.Nil(t, err)
// 		assert.Equal(t, "andri", "andri", resUserName, resToken)
// 	})

// 	t.Run("Test Login User Failed", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		var loginEmail = "andri@gmail.com"
// 		var loginPass = "qwerty"

// 		resUserName, resToken, err := userBusiness.LoginUser(loginEmail, loginPass)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "andri", "andri", resUserName, resToken)
// 	})
// }

// func TestUpdateDataUser(t *testing.T) {
// 	t.Run("Test Update Data User Succes", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserData{})
// 		editUser := users.Core{
// 			FullName:    "Andri G",
// 			UserName:    "Andri",
// 			Email:       "andri@gmail.com",
// 			Password:    "qwerty",
// 			PhoneNumber: "081234",
// 			Address:     "Jalan Urip",
// 			ImageURL:    "profile1.png",
// 		}
// 		result, err := userBusiness.UpdateDataUser(0, editUser)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, result)
// 	})
// 	t.Run("Test Update Data User Failed", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		editUser := users.Core{
// 			UserName:    "Andri",
// 			Email:       "andri@gmail.com",
// 			Password:    "qwerty",
// 			PhoneNumber: "081234",
// 			Address:     "Jalan Urip",
// 			ImageURL:    "profile1.png",
// 		}
// 		result, err := userBusiness.UpdateDataUser(0, editUser)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, -1, result)
// 	})
// }

// func TestSelectUser(t *testing.T) {
// 	t.Run("Test Select User Success", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserData{})
// 		result, err := userBusiness.SelectUser(0)
// 		assert.Nil(t, err)
// 		assert.Equal(t, "Andri G", result.FullName)
// 	})
// 	t.Run("Test Select User Failed", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		result, err := userBusiness.SelectUser(0)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, users.Core{}, result)
// 	})
// }

// func TestDeleteUser(t *testing.T) {
// 	t.Run("Test Delete User Succes", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserData{})
// 		result, err := userBusiness.DeleteUser(0)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, result)
// 	})
// 	t.Run("Test Delete User Failed", func(t *testing.T) {
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		result, err := userBusiness.DeleteUser(0)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 0, result)
// 	})
// }
