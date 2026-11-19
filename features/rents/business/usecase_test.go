package business

import (
	"be9/mnroom/features/rents"
	"be9/mnroom/helper"
	"be9/mnroom/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPaymentGateway struct {
	mock.Mock
}

func (m *MockPaymentGateway) ChargeBankTransfer(orderID string, amount int, bank string, customer helper.MidtransCustomer) (*helper.PaymentChargeResponse, error) {
	args := m.Called(orderID, amount, bank, customer)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*helper.PaymentChargeResponse), args.Error(1)
}

func (m *MockPaymentGateway) VerifySignature(orderID, statusCode, grossAmount, signatureKey string) bool {
	args := m.Called(orderID, statusCode, grossAmount, signatureKey)
	return args.Bool(0)
}

func (m *MockPaymentGateway) NormalizeStatus(transactionStatus, fraudStatus string) string {
	args := m.Called(transactionStatus, fraudStatus)
	return args.String(0)
}

func TestCreateRent(t *testing.T) {
	rentRepo := new(mocks.RentData)
	paymentGw := new(MockPaymentGateway)

	reqRent := rents.Core{
		DateStart: "2026-12-01",
		DateEnd:   "2026-12-03",
		Bank:      "bca",
		Room: rents.Rooms{
			ID: 1,
		},
	}

	t.Run("Success Create Rent", func(t *testing.T) {
		rentRepo.On("GetDataRentToken", mock.Anything, 10, 1).Return(0, nil).Once()
		rentRepo.On("GetDataRentUser", mock.Anything, 1, "2026-12-01", "2026-12-03").Return(0, nil).Once()
		rentRepo.On("GetData", mock.Anything, 1).Return(500000, nil).Once()
		rentRepo.On("InsertData", mock.Anything, mock.Anything).Return(100, nil).Once()
		rentRepo.On("GetDataUser", mock.Anything, 10).Return(rents.CoreUser{
			Username: "testuser",
			Email:    "test@mail.com",
			Phone:    "08123456789",
			Address:  "Jl. Test",
		}, nil).Once()

		paymentGw.On("ChargeBankTransfer", mock.Anything, 1000000, "bca", mock.Anything).Return(&helper.PaymentChargeResponse{
			TransactionID:     "trx-123",
			OrderID:           "RENT-100-12345",
			PaymentType:       "bank_transfer",
			BankTransfer:      "bca",
			GrossAmount:       1000000,
			VANumber:          "1234567890",
			TransactionStatus: "pending",
		}, nil).Once()

		rentRepo.On("InsertDataPayment", mock.Anything, mock.Anything).Return(rents.CorePayments{
			ID:                1,
			TransactionID:     "trx-123",
			OrderID:           "RENT-100-12345",
			PaymentType:       "bank_transfer",
			BankTransfer:      "bca",
			GrossAmount:       1000000,
			VANumber:          "1234567890",
			TransactionStatus: "pending",
		}, nil).Once()

		srv := NewRentBusiness(rentRepo, paymentGw)
		res, err := srv.CreateRent(context.Background(), 10, reqRent)

		assert.NoError(t, err)
		assert.Equal(t, "trx-123", res.TransactionID)
		assert.Equal(t, "1234567890", res.VANumber)
		rentRepo.AssertExpectations(t)
		paymentGw.AssertExpectations(t)
	})

	t.Run("Already Booked By User", func(t *testing.T) {
		rentRepo.On("GetDataRentToken", mock.Anything, 10, 1).Return(1, nil).Once()

		srv := NewRentBusiness(rentRepo, paymentGw)
		_, err := srv.CreateRent(context.Background(), 10, reqRent)

		assert.ErrorIs(t, err, rents.ErrAlreadyBooked)
		rentRepo.AssertExpectations(t)
	})

	t.Run("Room Already Booked For Dates", func(t *testing.T) {
		rentRepo.On("GetDataRentToken", mock.Anything, 10, 1).Return(0, nil).Once()
		rentRepo.On("GetDataRentUser", mock.Anything, 1, "2026-12-01", "2026-12-03").Return(1, nil).Once()

		srv := NewRentBusiness(rentRepo, paymentGw)
		_, err := srv.CreateRent(context.Background(), 10, reqRent)

		assert.ErrorIs(t, err, rents.ErrRoomBooked)
		rentRepo.AssertExpectations(t)
	})

	t.Run("Invalid Date Range", func(t *testing.T) {
		rentRepo.On("GetDataRentToken", mock.Anything, 10, 1).Return(0, nil).Once()
		rentRepo.On("GetDataRentUser", mock.Anything, 1, "2026-12-03", "2026-12-01").Return(0, nil).Once()
		rentRepo.On("GetData", mock.Anything, 1).Return(500000, nil).Once()

		invalidRent := reqRent
		invalidRent.DateStart = "2026-12-03"
		invalidRent.DateEnd = "2026-12-01"

		srv := NewRentBusiness(rentRepo, paymentGw)
		_, err := srv.CreateRent(context.Background(), 10, invalidRent)

		assert.ErrorIs(t, err, rents.ErrInvalidDate)
		rentRepo.AssertExpectations(t)
	})
}

func TestGetDataRent(t *testing.T) {
	rentRepo := new(mocks.RentData)
	paymentGw := new(MockPaymentGateway)

	t.Run("Success Get Data", func(t *testing.T) {
		rentRepo.On("GetDataRent", mock.Anything, 1).Return([]rents.Core{{ID: 1, Bank: "bca"}}, nil).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetDataRent(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(res))
		rentRepo.AssertExpectations(t)
	})

	t.Run("Error Get Data", func(t *testing.T) {
		rentRepo.On("GetDataRent", mock.Anything, 1).Return(nil, errors.New("db error")).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetDataRent(context.Background(), 1)
		assert.Error(t, err)
		assert.Nil(t, res)
		rentRepo.AssertExpectations(t)
	})
}
