package presentation

import (
	"be9/mnroom/features/feedback"
	"be9/mnroom/features/feedback/presentation/request"
	"be9/mnroom/features/feedback/presentation/response"
	"be9/mnroom/helper"
	_middlewares "be9/mnroom/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FeedbackHandler struct {
	feedbackBusiness feedback.Business
}

func NewFeedbackHandler(business feedback.Business) *FeedbackHandler {
	return &FeedbackHandler{
		feedbackBusiness: business,
	}
}

func (h *FeedbackHandler) InsertFeedback(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}

	var insertFeedback request.Feedback
	errBind := c.Bind(&insertFeedback)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}

	// data, _ := h.feedbackBusiness.GetDataRentUser(idToken, int(insertFeedback.RentsID))
	newFeedback := request.ToCore(insertFeedback)
	newFeedback.User.ID = idToken
	// newFeedback.Rents.ID = data
	row, err := h.feedbackBusiness.InsertFeedback(newFeedback)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert feedback"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert feedback"))
}

func (h *FeedbackHandler) GetDataRoom(c echo.Context) error {
	id := c.Param("id")
	idRoom, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid id"))
	}
	// data, _ := h.feedbackBusiness.GetDataRoom(idRoom)
	// dataRentInt, _ := h.feedbackBusiness.GetDataRent(data)
	dataFeedback, err := h.feedbackBusiness.GetFeedbackByRoom(idRoom)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(dataFeedback)))
}
