package presentation

import (
	"be9/mnroom/features/payments"
	"be9/mnroom/features/payments/presentation/response"
	"be9/mnroom/helper"
	"net/http"

	_middlewares "be9/mnroom/middlewares"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentBusiness payments.Business
}

func NewPaymentHandler(business payments.Business) *PaymentHandler {
	return &PaymentHandler{
		paymentBusiness: business,
	}
}

func (y *PaymentHandler) GetAllData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	data, err := y.paymentBusiness.GetAllData(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", response.FromCoreListPayments(data)))
}
