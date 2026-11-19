package business

import (
	"be9/mnroom/features/login"
	"be9/mnroom/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuth(t *testing.T) {
	repo := new(mocks.LoginData)
	insertData := login.Core{
		Email:    "mulya@mail.com",
		Password: "123456",
	}
	expectedData := login.Core{
		ID:       1,
		Email:    "mulya@mail.com",
		Password: "hashedpassword",
	}

	t.Run("Success Auth", func(t *testing.T) {
		repo.On("Auth", mock.Anything, mock.Anything).Return(expectedData, nil).Once()
		srv := NewAuthBusiness(repo)

		res, err := srv.Auth(context.Background(), insertData)
		assert.NoError(t, err)
		assert.Equal(t, expectedData.ID, res.ID)
		assert.Equal(t, expectedData.Email, res.Email)
		repo.AssertExpectations(t)
	})

	t.Run("Error Auth", func(t *testing.T) {
		repo.On("Auth", mock.Anything, mock.Anything).Return(login.Core{}, errors.New("there is some error")).Once()
		srv := NewAuthBusiness(repo)

		res, err := srv.Auth(context.Background(), insertData)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, login.Core{}, res)
		repo.AssertExpectations(t)
	})
}
