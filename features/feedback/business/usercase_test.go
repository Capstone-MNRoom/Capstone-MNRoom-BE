package business

import (
	"be9/mnroom/features/feedback"
	"be9/mnroom/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertFeedback(t *testing.T) {
	repo := new(mocks.FeedbackData)
	insertData := feedback.Core{
		Rating:  5,
		Comment: "Great place!",
		User: feedback.User{
			ID: 1,
		},
		Rents: feedback.Rents{
			ID: 1,
		},
	}

	t.Run("Success Insert Feedback", func(t *testing.T) {
		repo.On("InsertFeedback", mock.Anything, insertData).Return(1, nil).Once()
		srv := NewFeedbackBusiness(repo)

		res, err := srv.InsertFeedback(context.Background(), insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Insert Feedback", func(t *testing.T) {
		repo.On("InsertFeedback", mock.Anything, insertData).Return(0, errors.New("insert failed")).Once()
		srv := NewFeedbackBusiness(repo)

		res, err := srv.InsertFeedback(context.Background(), insertData)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}

func TestGetDataRoom(t *testing.T) {
	repo := new(mocks.FeedbackData)

	t.Run("Success Get Data Room", func(t *testing.T) {
		repo.On("GetDataRoom", mock.Anything, 1).Return(1, nil).Once()
		srv := NewFeedbackBusiness(repo)

		res, err := srv.GetDataRoom(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data Room", func(t *testing.T) {
		repo.On("GetDataRoom", mock.Anything, 1).Return(0, errors.New("not found")).Once()
		srv := NewFeedbackBusiness(repo)

		res, err := srv.GetDataRoom(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}

func TestGetDataRent(t *testing.T) {
	repo := new(mocks.FeedbackData)

	t.Run("Success Get Data Rent", func(t *testing.T) {
		repo.On("GetDataRent", mock.Anything, 1).Return(1, nil).Once()
		srv := NewFeedbackBusiness(repo)

		res, err := srv.GetDataRent(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data Rent", func(t *testing.T) {
		repo.On("GetDataRent", mock.Anything, 1).Return(0, errors.New("not found")).Once()
		srv := NewFeedbackBusiness(repo)

		res, err := srv.GetDataRent(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}

func TestGetFeedbackByRoom(t *testing.T) {
	repo := new(mocks.FeedbackData)
	feedbackList := []feedback.Core{
		{
			ID:      1,
			Rating:  5,
			Comment: "Awesome room!",
			User:    feedback.User{ID: 1, Username: "John"},
		},
	}

	t.Run("Success Get Feedback By Room", func(t *testing.T) {
		repo.On("GetFeedbackByRoom", mock.Anything, 1).Return(feedbackList, nil).Once()
		srv := NewFeedbackBusiness(repo)

		res, err := srv.GetFeedbackByRoom(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, 5, res[0].Rating)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Feedback By Room", func(t *testing.T) {
		repo.On("GetFeedbackByRoom", mock.Anything, 1).Return(nil, errors.New("query error")).Once()
		srv := NewFeedbackBusiness(repo)

		res, err := srv.GetFeedbackByRoom(context.Background(), 1)
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetDataRentUser(t *testing.T) {
	repo := new(mocks.FeedbackData)

	t.Run("Success Get Data Rent User", func(t *testing.T) {
		repo.On("GetDataRentUser", mock.Anything, 1, 2).Return(1, nil).Once()
		srv := NewFeedbackBusiness(repo)

		res, err := srv.GetDataRentUser(context.Background(), 1, 2)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data Rent User", func(t *testing.T) {
		repo.On("GetDataRentUser", mock.Anything, 1, 2).Return(0, errors.New("query error")).Once()
		srv := NewFeedbackBusiness(repo)

		res, err := srv.GetDataRentUser(context.Background(), 1, 2)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}
