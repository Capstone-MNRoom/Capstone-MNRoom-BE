package presentation

import (
	"be9/mnroom/features/payments"
	_payments "be9/mnroom/features/payments/presentation/request"
	_response "be9/mnroom/features/payments/presentation/response"
	"be9/mnroom/features/rents"
	"be9/mnroom/features/rents/presentation/request"
	"be9/mnroom/features/rents/presentation/response"
	"be9/mnroom/helper"
	_middlewares "be9/mnroom/middlewares"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RentHandler struct {
	rentBusiness rents.Business
}

func NewEventHandler(business rents.Business) *RentHandler {
	return &RentHandler{
		rentBusiness: business,
	}
}

func (t *RentHandler) GetData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	var insertData request.Rents
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	v := validator.New()
	errDateStart := v.Var(insertData.DateStart, "required")
	if errDateStart != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("date start cannot be empty"))
	}
	errDateEnd := v.Var(insertData.DateEnd, "required")
	if errDateEnd != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("date end cannot be empty"))
	}
	errBank := v.Var(insertData.Bank, "required")
	if errBank != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("bank cannot be empty"))
	}
	newRent := request.ToCore(insertData)
	rowToken, _ := t.rentBusiness.GetDataRentToken(idToken)
	if rowToken != 1 {
		rowDataRent, _ := t.rentBusiness.GetDataRentUser(newRent.Room.ID, newRent.DateStart, newRent.DateEnd)
		if rowDataRent != 1 {
			data, _ := t.rentBusiness.GetData(newRent.Room.ID)
			date1, _ := time.Parse("2006-01-02", newRent.DateStart)
			date2, _ := time.Parse("2006-01-02", newRent.DateEnd)
			difference := date2.Sub(date1)
			newRent.TotalRentalPrice = data * int(int64(difference.Hours()/24))
			newRent.Status = "Success"
			newRent.User.ID = idToken
			rowIDRent, err := t.rentBusiness.InsertData(newRent)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
			}
			rand.Seed(5)
			var insertPayment _payments.Payments
			newPayment := _payments.ToCore(insertPayment)
			newPayment.TransactionID = "be03df7d-2f97-4c8c-a53c-8959f1b67295"
			newPayment.PaymentType = "bank transfer"
			newPayment.OrderID = rand.Int()
			newPayment.BankTransfer = newRent.Bank
			newPayment.GrossAmount = newRent.TotalRentalPrice
			newPayment.VANumber = 812785002530231
			newPayment.TransactionStatus = "Pending"
			newPayment.Rents.ID = rowIDRent
			newPayment.User.ID = idToken
			dataPayment, _ := t.rentBusiness.InsertDataPayment(rents.CorePayments(newPayment))
			return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to insert data", _response.FromCore(payments.Core(dataPayment))))
		}
		return c.JSON(http.StatusMethodNotAllowed, helper.ResponseFailed("this room has booked"))
	}
	return c.JSON(http.StatusMethodNotAllowed, helper.ResponseFailed("you already booked this room"))
}

func (t *RentHandler) GetDataRent(c echo.Context) error {
	id := c.Param("id")
	idRoom, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid id"))
	}
	data, err := t.rentBusiness.GetDataRent(idRoom)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}
