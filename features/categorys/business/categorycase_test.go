package business

import (
	"be9/mnroom/features/categorys"
	"be9/mnroom/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertData(t *testing.T) {
	repo := new(mocks.CategoryData)
	insertData := categorys.Core{
		CategoryName: "Meeting",
	}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("InsertData", mock.Anything, insertData).Return(1, nil).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.InsertData(context.Background(), insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Insert", func(t *testing.T) {
		repo.On("InsertData", mock.Anything, insertData).Return(0, errors.New("insert failed")).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.InsertData(context.Background(), insertData)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}

func TestGetDataAll(t *testing.T) {
	repo := new(mocks.CategoryData)
	categoryList := []categorys.Core{
		{ID: 1, CategoryName: "Meeting Room"},
		{ID: 2, CategoryName: "Coworking Space"},
	}

	t.Run("Success Get Data All", func(t *testing.T) {
		repo.On("GetDataAll", mock.Anything).Return(categoryList, nil).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.GetDataAll(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 2, len(res))
		assert.Equal(t, "Meeting Room", res[0].CategoryName)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data All", func(t *testing.T) {
		repo.On("GetDataAll", mock.Anything).Return(nil, errors.New("query error")).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.GetDataAll(context.Background())
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetData(t *testing.T) {
	repo := new(mocks.CategoryData)
	categoryData := categorys.Core{
		ID:           1,
		CategoryName: "Meeting Room",
	}

	t.Run("Success Get Data", func(t *testing.T) {
		repo.On("GetData", mock.Anything, 1).Return(categoryData, nil).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.GetData(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, categoryData.ID, res.ID)
		assert.Equal(t, categoryData.CategoryName, res.CategoryName)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data", func(t *testing.T) {
		repo.On("GetData", mock.Anything, 1).Return(categorys.Core{}, errors.New("not found")).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.GetData(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, categorys.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateData(t *testing.T) {
	repo := new(mocks.CategoryData)
	updateData := categorys.Core{
		CategoryName: "Updated Meeting Room",
	}

	t.Run("Success Update", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, 1, updateData).Return(1, nil).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.UpdateData(context.Background(), 1, updateData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Update", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, 1, updateData).Return(0, errors.New("update failed")).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.UpdateData(context.Background(), 1, updateData)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T) {
	repo := new(mocks.CategoryData)

	t.Run("Success Delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, 1).Return(1, nil).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.DeleteData(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, 1).Return(0, errors.New("delete failed")).Once()
		srv := NewCategoryBusiness(repo)

		res, err := srv.DeleteData(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}
