package business

import (
	"be9/mnroom/features/rents"
	"be9/mnroom/features/users"
	"be9/mnroom/helper"
	"context"
	"fmt"
	"time"
)

type rentUseCase struct {
	rentData       rents.Data
	paymentGateway helper.PaymentGateway
}

func NewRentBusiness(rent rents.Data, payment helper.PaymentGateway) rents.Business {
	return &rentUseCase{
		rentData:       rent,
		paymentGateway: payment,
	}
}

func (rnt *rentUseCase) CreateRent(ctx context.Context, idToken int, newRent rents.Core) (dataPayment rents.CorePayments, err error) {
	// 1. Check if user already booked this room
	rowToken, _ := rnt.rentData.GetDataRentToken(ctx, idToken, newRent.Room.ID)
	if rowToken == 1 {
		return rents.CorePayments{}, rents.ErrAlreadyBooked
	}

	// 2. Check if room is already booked by others for specified dates
	rowDataRent, _ := rnt.rentData.GetDataRentUser(ctx, newRent.Room.ID, newRent.DateStart, newRent.DateEnd)
	if rowDataRent == 1 {
		return rents.CorePayments{}, rents.ErrRoomBooked
	}

	// 3. Fetch room price per day
	roomPrice, err := rnt.rentData.GetData(ctx, newRent.Room.ID)
	if err != nil {
		return rents.CorePayments{}, err
	}

	// 4. Calculate duration in days & total rental price
	date1, errStart := time.Parse("2006-01-02", newRent.DateStart)
	date2, errEnd := time.Parse("2006-01-02", newRent.DateEnd)
	if errStart != nil || errEnd != nil || !date2.After(date1) {
		return rents.CorePayments{}, rents.ErrInvalidDate
	}
	days := int(date2.Sub(date1).Hours() / 24)
	if days < 1 {
		days = 1
	}

	newRent.TotalRentalPrice = roomPrice * days
	newRent.Status = "Pending"
	newRent.User.ID = idToken

	// 5. Insert rent transaction
	rowIDRent, err := rnt.rentData.InsertData(ctx, newRent)
	if err != nil {
		return rents.CorePayments{}, err
	}

	// 6. Fetch user details for payment customer data
	userData, err := rnt.rentData.GetDataUser(ctx, idToken)
	if err != nil {
		return rents.CorePayments{}, err
	}

	// 7. Request payment charge to Payment Gateway (Midtrans)
	orderID := fmt.Sprintf("RENT-%d-%d", rowIDRent, time.Now().Unix())
	customer := helper.MidtransCustomer{
		Username: userData.Username,
		Phone:    userData.Phone,
		Email:    userData.Email,
		Address:  userData.Address,
	}

	chargeRes, errCharge := rnt.paymentGateway.ChargeBankTransfer(orderID, newRent.TotalRentalPrice, newRent.Bank, customer)
	if errCharge != nil {
		return rents.CorePayments{}, errCharge
	}

	// 8. Save payment record
	newPayment := rents.CorePayments{
		TransactionID:     chargeRes.TransactionID,
		OrderID:           chargeRes.OrderID,
		PaymentType:       chargeRes.PaymentType,
		BankTransfer:      newRent.Bank,
		GrossAmount:       chargeRes.GrossAmount,
		VANumber:          chargeRes.VANumber,
		TransactionStatus: chargeRes.TransactionStatus,
		Rents:             rents.Core{ID: rowIDRent},
		User:              users.Core{ID: idToken},
	}

	dataPayment, err = rnt.rentData.InsertDataPayment(ctx, newPayment)
	if err != nil {
		return rents.CorePayments{}, err
	}

	return dataPayment, nil
}

func (rnt *rentUseCase) GetData(ctx context.Context, id int) (data int, err error) {
	data, err = rnt.rentData.GetData(ctx, id)
	return data, err
}

func (rnt *rentUseCase) InsertData(ctx context.Context, insert rents.Core) (row int, err error) {
	row, err = rnt.rentData.InsertData(ctx, insert)
	return row, err
}

func (rnt *rentUseCase) GetDataRentToken(ctx context.Context, idToken int, idRoom int) (row int, err error) {
	row, err = rnt.rentData.GetDataRentToken(ctx, idToken, idRoom)
	return row, err
}

func (rnt *rentUseCase) GetDataRentUser(ctx context.Context, id int, start string, end string) (row int, err error) {
	row, err = rnt.rentData.GetDataRentUser(ctx, id, start, end)
	return row, err
}

func (rnt *rentUseCase) GetDataRent(ctx context.Context, id int) (data []rents.Core, err error) {
	data, err = rnt.rentData.GetDataRent(ctx, id)
	return data, err
}

func (rnt *rentUseCase) InsertDataPayment(ctx context.Context, insert rents.CorePayments) (data rents.CorePayments, err error) {
	data, err = rnt.rentData.InsertDataPayment(ctx, insert)
	return data, err
}

func (rnt *rentUseCase) GetDataUser(ctx context.Context, idToken int) (data rents.CoreUser, err error) {
	data, err = rnt.rentData.GetDataUser(ctx, idToken)
	return data, err
}

func (rnt *rentUseCase) GetDataRentUserHistory(ctx context.Context, idToken int) (data []rents.Core, err error) {
	data, err = rnt.rentData.GetDataRentUserHistory(ctx, idToken)
	return data, err
}
