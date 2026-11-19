package business

import (
	"be9/mnroom/features/users"
	"be9/mnroom/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertData(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := users.Core{
		Username: "alta",
		Email:    "alta@mail.id",
		Password: "qwerty",
		Phone:    "0899123456",
		Address:  "Jl.Rambutan"}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("InsertData", mock.Anything, mock.Anything).Return(1, nil).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.InsertData(context.Background(), insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		repo.On("InsertData", mock.Anything, mock.Anything).Return(0, errors.New("there is some error")).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.InsertData(context.Background(), insertData)
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

		res, err := srv.GetAllData(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get All", func(t *testing.T) {
		repo.On("GetAllData", mock.Anything).Return(nil, errors.New("data not found")).Once()

		srv := NewUserBusiness(repo)

		res, err := srv.GetAllData(context.Background())
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetData(t *testing.T) {
	repo := new(mocks.UserData)
	userData := users.Core{
		ID:       1,
		Username: "Mulya",
		Email:    "mulya@mail.com",
		Phone:    "08995555",
		Address:  "Jl.Rambutan",
	}

	t.Run("Success Get Data", func(t *testing.T) {
		repo.On("GetData", mock.Anything, 1).Return(userData, nil).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.GetData(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, userData.ID, res.ID)
		assert.Equal(t, userData.Username, res.Username)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data", func(t *testing.T) {
		repo.On("GetData", mock.Anything, 1).Return(users.Core{}, errors.New("record not found")).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.GetData(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, users.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateData(t *testing.T) {
	repo := new(mocks.UserData)
	updateData := users.Core{
		Username: "Mulya Updated",
		Phone:    "08999999",
	}

	t.Run("Success Update", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, 1, updateData).Return(1, nil).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.UpdateData(context.Background(), 1, updateData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Update", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, 1, updateData).Return(0, errors.New("update failed")).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.UpdateData(context.Background(), 1, updateData)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Success Delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, 1).Return(1, nil).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.DeleteData(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, 1).Return(0, errors.New("delete failed")).Once()
		srv := NewUserBusiness(repo)

		res, err := srv.DeleteData(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}
