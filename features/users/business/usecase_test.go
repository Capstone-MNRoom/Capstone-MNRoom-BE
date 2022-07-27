package business

import (
	"be9/mnroom/features/users"
	"be9/mnroom/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertData(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := users.Core{ID: 1,
		Username: "alta",
		Email:    "alta@mail.id",
		Password: "qwerty",

		Address: "Jl.Rambutan"}
	// returnData := users.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Test InsertData", mock.Anything).Return(1, nil).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.InsertData(insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		repo.On("Test InsertData", mock.Anything).Return(0, errors.New("there is some error")).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.InsertData(insertData)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})

	// t.Run("Error insert when incomplete data", func(t *testing.T) {
	// 	/*
	// 		dont need to write repo.On because this test case dont need to call data layer. just handle on business layer.
	// 	*/
	// 	// repo.On("InsertData", mock.Anything).Return(-1, errors.New("all input data must be filled")).Once()
	// 	srv := NewUserBusiness(repo)

	// 	_, err := srv.InsertData(users.Core{})
	// 	assert.EqualError(t, err, "all input data must be filled")
	// 	repo.AssertExpectations(t)
	// })
}

// func TestGetAllData(t *testing.T) {
// 	repo := new(mocks.UserData)
// 	returnData := []users.Core{{ID: 1, Username: "Mulya", Email: "mulya@mail.com", Password: "123456", Phone: "08995555", Address: "Jl.Rambutan", Image: "https://storage.googleapis.com/event2022/Logo.png"}}

// 	t.Run("Success Get All", func(t *testing.T) {
// 		repo.On("GetAllData", mock.Anything).Return(returnData, nil).Once()

// 		srv := NewUserBusiness(repo)

// 		res, err := srv.GetAllData()
// 		assert.NoError(t, err)
// 		assert.Equal(t, returnData[0].ID, res[0].ID)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Error Get All", func(t *testing.T) {
// 		repo.On("SelectData", "").Return(nil, errors.New("data not found")).Once()

// 		srv := NewUserBusiness(repo)

// 		res, err := srv.GetAllData()
// 		assert.Error(t, err)
// 		assert.Nil(t, res)
// 		repo.AssertExpectations(t)
// 	})
// }
