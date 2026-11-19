package business

import (
	"be9/mnroom/features/rents"
	"be9/mnroom/features/rooms"
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
	reqRent := rents.Core{
		DateStart: "2026-12-01",
		DateEnd:   "2026-12-03",
		Bank:      "bca",
		Room: rooms.Core{
			ID: 1,
		},
	}

	t.Run("Success Create Rent", func(t *testing.T) {
		rentRepo := new(mocks.RentData)
		paymentGw := new(MockPaymentGateway)
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
		rentRepo := new(mocks.RentData)
		paymentGw := new(MockPaymentGateway)
		rentRepo.On("GetDataRentToken", mock.Anything, 10, 1).Return(1, nil).Once()

		srv := NewRentBusiness(rentRepo, paymentGw)
		_, err := srv.CreateRent(context.Background(), 10, reqRent)

		assert.ErrorIs(t, err, rents.ErrAlreadyBooked)
		rentRepo.AssertExpectations(t)
	})

	t.Run("Room Already Booked For Dates", func(t *testing.T) {
		rentRepo := new(mocks.RentData)
		paymentGw := new(MockPaymentGateway)
		rentRepo.On("GetDataRentToken", mock.Anything, 10, 1).Return(0, nil).Once()
		rentRepo.On("GetDataRentUser", mock.Anything, 1, "2026-12-01", "2026-12-03").Return(1, nil).Once()

		srv := NewRentBusiness(rentRepo, paymentGw)
		_, err := srv.CreateRent(context.Background(), 10, reqRent)

		assert.ErrorIs(t, err, rents.ErrRoomBooked)
		rentRepo.AssertExpectations(t)
	})

	t.Run("Invalid Date Range", func(t *testing.T) {
		rentRepo := new(mocks.RentData)
		paymentGw := new(MockPaymentGateway)
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

	t.Run("Error Fetch Room Price", func(t *testing.T) {
		rentRepo := new(mocks.RentData)
		paymentGw := new(MockPaymentGateway)
		rentRepo.On("GetDataRentToken", mock.Anything, 10, 1).Return(0, nil).Once()
		rentRepo.On("GetDataRentUser", mock.Anything, 1, "2026-12-01", "2026-12-03").Return(0, nil).Once()
		rentRepo.On("GetData", mock.Anything, 1).Return(0, errors.New("db error")).Once()

		srv := NewRentBusiness(rentRepo, paymentGw)
		_, err := srv.CreateRent(context.Background(), 10, reqRent)

		assert.Error(t, err)
		rentRepo.AssertExpectations(t)
	})

	t.Run("Error Insert Rent", func(t *testing.T) {
		rentRepo := new(mocks.RentData)
		paymentGw := new(MockPaymentGateway)
		rentRepo.On("GetDataRentToken", mock.Anything, 10, 1).Return(0, nil).Once()
		rentRepo.On("GetDataRentUser", mock.Anything, 1, "2026-12-01", "2026-12-03").Return(0, nil).Once()
		rentRepo.On("GetData", mock.Anything, 1).Return(500000, nil).Once()
		rentRepo.On("InsertData", mock.Anything, mock.Anything).Return(0, errors.New("insert rent failed")).Once()

		srv := NewRentBusiness(rentRepo, paymentGw)
		_, err := srv.CreateRent(context.Background(), 10, reqRent)

		assert.Error(t, err)
		assert.EqualError(t, err, "insert rent failed")
		rentRepo.AssertExpectations(t)
	})

	t.Run("Error Payment Gateway Charge", func(t *testing.T) {
		rentRepo := new(mocks.RentData)
		paymentGw := new(MockPaymentGateway)
		rentRepo.On("GetDataRentToken", mock.Anything, 10, 1).Return(0, nil).Once()
		rentRepo.On("GetDataRentUser", mock.Anything, 1, "2026-12-01", "2026-12-03").Return(0, nil).Once()
		rentRepo.On("GetData", mock.Anything, 1).Return(500000, nil).Once()
		rentRepo.On("InsertData", mock.Anything, mock.Anything).Return(100, nil).Once()
		rentRepo.On("GetDataUser", mock.Anything, 10).Return(rents.CoreUser{
			Username: "testuser",
			Email:    "test@mail.com",
			Phone:    "08123456789",
		}, nil).Once()
		paymentGw.On("ChargeBankTransfer", mock.Anything, 1000000, "bca", mock.Anything).Return(nil, errors.New("midtrans connection error")).Once()

		srv := NewRentBusiness(rentRepo, paymentGw)
		_, err := srv.CreateRent(context.Background(), 10, reqRent)

		assert.Error(t, err)
		assert.EqualError(t, err, "midtrans connection error")
		rentRepo.AssertExpectations(t)
		paymentGw.AssertExpectations(t)
	})

	t.Run("Error Insert Payment", func(t *testing.T) {
		rentRepo := new(mocks.RentData)
		paymentGw := new(MockPaymentGateway)
		rentRepo.On("GetDataRentToken", mock.Anything, 10, 1).Return(0, nil).Once()
		rentRepo.On("GetDataRentUser", mock.Anything, 1, "2026-12-01", "2026-12-03").Return(0, nil).Once()
		rentRepo.On("GetData", mock.Anything, 1).Return(500000, nil).Once()
		rentRepo.On("InsertData", mock.Anything, mock.Anything).Return(100, nil).Once()
		rentRepo.On("GetDataUser", mock.Anything, 10).Return(rents.CoreUser{
			Username: "testuser",
			Email:    "test@mail.com",
			Phone:    "08123456789",
		}, nil).Once()
		paymentGw.On("ChargeBankTransfer", mock.Anything, 1000000, "bca", mock.Anything).Return(&helper.PaymentChargeResponse{
			TransactionID:     "midtrans-txn-123",
			OrderID:           "ORDER-123",
			GrossAmount:       1000000,
			PaymentType:       "bank_transfer",
			BankTransfer:      "bca",
			VANumber:          "12345678901",
			TransactionStatus: "pending",
		}, nil).Once()
		rentRepo.On("InsertDataPayment", mock.Anything, mock.Anything).Return(rents.CorePayments{}, errors.New("insert payment failed")).Once()

		srv := NewRentBusiness(rentRepo, paymentGw)
		_, err := srv.CreateRent(context.Background(), 10, reqRent)

		assert.Error(t, err)
		assert.EqualError(t, err, "insert payment failed")
		rentRepo.AssertExpectations(t)
		paymentGw.AssertExpectations(t)
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

func TestGetData(t *testing.T) {
	rentRepo := new(mocks.RentData)
	paymentGw := new(MockPaymentGateway)

	t.Run("Success Get Data", func(t *testing.T) {
		rentRepo.On("GetData", mock.Anything, 1).Return(500000, nil).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetData(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 500000, res)
		rentRepo.AssertExpectations(t)
	})

	t.Run("Error Get Data", func(t *testing.T) {
		rentRepo.On("GetData", mock.Anything, 1).Return(0, errors.New("not found")).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetData(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		rentRepo.AssertExpectations(t)
	})
}

func TestInsertData(t *testing.T) {
	rentRepo := new(mocks.RentData)
	paymentGw := new(MockPaymentGateway)
	newRent := rents.Core{
		DateStart: "2026-12-01",
		DateEnd:   "2026-12-03",
		Bank:      "bca",
	}

	t.Run("Success Insert Data", func(t *testing.T) {
		rentRepo.On("InsertData", mock.Anything, newRent).Return(1, nil).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.InsertData(context.Background(), newRent)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		rentRepo.AssertExpectations(t)
	})

	t.Run("Error Insert Data", func(t *testing.T) {
		rentRepo.On("InsertData", mock.Anything, newRent).Return(0, errors.New("insert failed")).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.InsertData(context.Background(), newRent)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		rentRepo.AssertExpectations(t)
	})
}

func TestGetDataRentToken(t *testing.T) {
	rentRepo := new(mocks.RentData)
	paymentGw := new(MockPaymentGateway)

	t.Run("Success Get Data Rent Token", func(t *testing.T) {
		rentRepo.On("GetDataRentToken", mock.Anything, 1, 2).Return(1, nil).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetDataRentToken(context.Background(), 1, 2)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		rentRepo.AssertExpectations(t)
	})

	t.Run("Error Get Data Rent Token", func(t *testing.T) {
		rentRepo.On("GetDataRentToken", mock.Anything, 1, 2).Return(0, errors.New("query error")).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetDataRentToken(context.Background(), 1, 2)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		rentRepo.AssertExpectations(t)
	})
}

func TestGetDataRentUser(t *testing.T) {
	rentRepo := new(mocks.RentData)
	paymentGw := new(MockPaymentGateway)

	t.Run("Success Get Data Rent User", func(t *testing.T) {
		rentRepo.On("GetDataRentUser", mock.Anything, 1, "2026-12-01", "2026-12-03").Return(1, nil).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetDataRentUser(context.Background(), 1, "2026-12-01", "2026-12-03")
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		rentRepo.AssertExpectations(t)
	})

	t.Run("Error Get Data Rent User", func(t *testing.T) {
		rentRepo.On("GetDataRentUser", mock.Anything, 1, "2026-12-01", "2026-12-03").Return(0, errors.New("query error")).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetDataRentUser(context.Background(), 1, "2026-12-01", "2026-12-03")
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		rentRepo.AssertExpectations(t)
	})
}

func TestInsertDataPayment(t *testing.T) {
	rentRepo := new(mocks.RentData)
	paymentGw := new(MockPaymentGateway)
	paymentData := rents.CorePayments{
		OrderID:     "ORDER-123",
		GrossAmount: 1000000,
	}

	t.Run("Success Insert Data Payment", func(t *testing.T) {
		rentRepo.On("InsertDataPayment", mock.Anything, paymentData).Return(paymentData, nil).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.InsertDataPayment(context.Background(), paymentData)
		assert.NoError(t, err)
		assert.Equal(t, paymentData.OrderID, res.OrderID)
		rentRepo.AssertExpectations(t)
	})

	t.Run("Error Insert Data Payment", func(t *testing.T) {
		rentRepo.On("InsertDataPayment", mock.Anything, paymentData).Return(rents.CorePayments{}, errors.New("db error")).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.InsertDataPayment(context.Background(), paymentData)
		assert.Error(t, err)
		assert.Equal(t, rents.CorePayments{}, res)
		rentRepo.AssertExpectations(t)
	})
}

func TestGetDataUser(t *testing.T) {
	rentRepo := new(mocks.RentData)
	paymentGw := new(MockPaymentGateway)
	userData := rents.CoreUser{
		ID:       1,
		Username: "testuser",
		Email:    "test@mail.com",
	}

	t.Run("Success Get Data User", func(t *testing.T) {
		rentRepo.On("GetDataUser", mock.Anything, 1).Return(userData, nil).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetDataUser(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, userData.Username, res.Username)
		rentRepo.AssertExpectations(t)
	})

	t.Run("Error Get Data User", func(t *testing.T) {
		rentRepo.On("GetDataUser", mock.Anything, 1).Return(rents.CoreUser{}, errors.New("not found")).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetDataUser(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, rents.CoreUser{}, res)
		rentRepo.AssertExpectations(t)
	})
}

func TestGetDataRentUserHistory(t *testing.T) {
	rentRepo := new(mocks.RentData)
	paymentGw := new(MockPaymentGateway)
	history := []rents.Core{
		{ID: 1, Bank: "bca", TotalRentalPrice: 500000},
	}

	t.Run("Success Get History", func(t *testing.T) {
		rentRepo.On("GetDataRentUserHistory", mock.Anything, 1).Return(history, nil).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetDataRentUserHistory(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(res))
		rentRepo.AssertExpectations(t)
	})

	t.Run("Error Get History", func(t *testing.T) {
		rentRepo.On("GetDataRentUserHistory", mock.Anything, 1).Return(nil, errors.New("query error")).Once()
		srv := NewRentBusiness(rentRepo, paymentGw)

		res, err := srv.GetDataRentUserHistory(context.Background(), 1)
		assert.Error(t, err)
		assert.Nil(t, res)
		rentRepo.AssertExpectations(t)
	})
}

