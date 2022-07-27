package business

import (
	"be9/mnroom/features/login"
	"be9/mnroom/mocks"
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

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("Auth", mock.Anything).Return(1, nil).Once()
		srv := NewAuthBusiness(repo)

		res, err := srv.Auth(insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		repo.On("Auth", mock.Anything).Return(0, errors.New("there is some error")).Once()
		srv := NewAuthBusiness(repo)

		res, err := srv.Auth(insertData)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}
