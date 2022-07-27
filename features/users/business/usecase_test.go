package business

import (
	"be9/mnroom/features/users"
	"be9/mnroom/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuth(t *testing.T) {
	repo := new(mocks.UserData)
<<<<<<< HEAD
	insertData := users.Core{
		Username: "alta",
		Email:    "alta@mail.id",
		Password: "qwerty",
		Phone:    "0899123456",
		Address:  "Jl.Rambutan"}
=======
	insertData := users.Core{ID: 1,
		Username: "alta",
		Email:    "alta@mail.id",
		Password: "qwerty",

		Address: "Jl.Rambutan"}
	// returnData := users.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}
>>>>>>> 2fb0074bcf5d8d24718673c96b9ea7897b7ae55b

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Auth", mock.Anything).Return(1, nil).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.InsertData(insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(0, errors.New("there is some error")).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.InsertData(insertData)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}

func TestGetAllData(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := []users.Core{{ID: 1, Username: "Mulya", Email: "mulya@mail.com", Password: "123456", Phone: "08995555", Address: "Jl.Rambutan", Image: "https://storage.googleapis.com/event2022/Logo.png"}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAllData", mock.Anything).Return(returnData, nil).Once()

		srv := NewUserBusiness(repo)

		res, err := srv.GetAllData()
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		repo.On("GetAllData", mock.Anything).Return(nil, errors.New("data not found")).Once()

		srv := NewUserBusiness(repo)

		res, err := srv.GetAllData()
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}
