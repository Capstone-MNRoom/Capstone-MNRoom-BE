package presentation

import (
	"be9/mnroom/features/payments"
	_response "be9/mnroom/features/payments/presentation/response"
	"be9/mnroom/helper"
	"encoding/json"
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
	data, err := y.paymentBusiness.GetAllData(c.Request().Context(), idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", _response.FromCoreListPayments(data)))
}

func (y *PaymentHandler) UpdateData(c echo.Context) error {
	var notificationPayload map[string]interface{}
	errDecode := json.NewDecoder(c.Request().Body).Decode(&notificationPayload)
	if errDecode != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to decode notification payload"))
	}

	orderId, _ := notificationPayload["order_id"].(string)
	transactionStatus, _ := notificationPayload["transaction_status"].(string)
	fraudStatus, _ := notificationPayload["fraud_status"].(string)

	if orderId == "" {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid order_id"))
	}

	// Normalize status using standard Midtrans rules
	gateway := helper.NewMidtransGateway()
	finalStatus := gateway.NormalizeStatus(transactionStatus, fraudStatus)

	row, err := y.paymentBusiness.UpdateData(c.Request().Context(), orderId, finalStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	if row == 0 {
		return c.JSON(http.StatusNotFound, helper.ResponseFailed("payment record not found"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to update payment status"))
}
