package presentation

import (
	"be9/mnroom/features/rents"
	"be9/mnroom/features/rents/presentation/request"
	"be9/mnroom/features/rents/presentation/response"
	"be9/mnroom/helper"
	_middlewares "be9/mnroom/middlewares"
	"net/http"
	"strconv"

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
			newRent.TotalRentalPrice = data
			newRent.Status = "Success"
			newRent.User.ID = idToken
			row, err := t.rentBusiness.InsertData(newRent)
			if row != 1 {
				return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert data"))
			}
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
			}
			return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert data"))
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
