package presentation

import (
	"be9/mnroom/features/payments"
	_response "be9/mnroom/features/payments/presentation/response"
	"be9/mnroom/features/rents"
	"be9/mnroom/features/rents/presentation/request"
	"be9/mnroom/features/rents/presentation/response"
	"be9/mnroom/helper"
	_middlewares "be9/mnroom/middlewares"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
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

func (t *RentHandler) CreateRent(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	var insertData request.Rents
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}

	v := validator.New()
	if err := v.Var(insertData.DateStart, "required"); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("date start cannot be empty"))
	}
	if err := v.Var(insertData.DateEnd, "required"); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("date end cannot be empty"))
	}
	if err := v.Var(insertData.RoomID, "required,numeric"); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("rooms cannot be empty"))
	}
	if err := v.Var(insertData.Bank, "required"); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("bank cannot be empty"))
	}

	newRent := request.ToCore(insertData)
	dataPayment, err := t.rentBusiness.CreateRent(c.Request().Context(), idToken, newRent)
	if err != nil {
		if errors.Is(err, rents.ErrAlreadyBooked) || errors.Is(err, rents.ErrRoomBooked) {
			return c.JSON(http.StatusMethodNotAllowed, helper.ResponseFailed(err.Error()))
		}
		if errors.Is(err, rents.ErrInvalidDate) {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to insert data", _response.FromCore(payments.Core(dataPayment))))
}

// GetData is kept for backward-compatibility with existing routes
func (t *RentHandler) GetData(c echo.Context) error {
	return t.CreateRent(c)
}

func (t *RentHandler) GetDataRent(c echo.Context) error {
	id := c.Param("id")
	idRoom, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid id"))
	}
	data, err := t.rentBusiness.GetDataRent(c.Request().Context(), idRoom)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}

func (t *RentHandler) GetDataRentUserHistory(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	data, err := t.rentBusiness.GetDataRentUserHistory(c.Request().Context(), idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}
