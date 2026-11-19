package business

import (
	"be9/mnroom/features/payments"
	"be9/mnroom/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllData(t *testing.T) {
	repo := new(mocks.PaymentData)
	paymentList := []payments.Core{
		{
			ID:                1,
			OrderID:           "ORDER-12345",
			GrossAmount:       1000000,
			TransactionStatus: "settlement",
		},
	}

	t.Run("Success Get All Data", func(t *testing.T) {
		repo.On("GetAllData", mock.Anything, 10).Return(paymentList, nil).Once()
		srv := NewPaymentBusiness(repo)

		res, err := srv.GetAllData(context.Background(), 10)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, "ORDER-12345", res[0].OrderID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get All Data", func(t *testing.T) {
		repo.On("GetAllData", mock.Anything, 10).Return(nil, errors.New("query error")).Once()
		srv := NewPaymentBusiness(repo)

		res, err := srv.GetAllData(context.Background(), 10)
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateData(t *testing.T) {
	repo := new(mocks.PaymentData)

	t.Run("Success Update Data", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, "ORDER-12345", "settlement").Return(1, nil).Once()
		srv := NewPaymentBusiness(repo)

		res, err := srv.UpdateData(context.Background(), "ORDER-12345", "settlement")
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Update Data", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, "ORDER-12345", "expire").Return(0, errors.New("record not found")).Once()
		srv := NewPaymentBusiness(repo)

		res, err := srv.UpdateData(context.Background(), "ORDER-12345", "expire")
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}
