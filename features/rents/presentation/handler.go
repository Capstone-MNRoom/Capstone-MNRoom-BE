package presentation

import (
	"be9/mnroom/features/rents"
	"be9/mnroom/features/rents/presentation/request"
	"be9/mnroom/helper"
	_middlewares "be9/mnroom/middlewares"
	"net/http"

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
	errValidator := v.Struct(insertData)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	newRent := request.ToCore(insertData)
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
