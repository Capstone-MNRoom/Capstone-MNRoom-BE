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

	newFeedback := request.ToCore(insertFeedback)
	newFeedback.User.ID = idToken
	row, err := h.feedbackBusiness.InsertFeedback(newFeedback)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert feedback"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert feedback"))
}

func (h *FeedbackHandler) GetFeedbackByRoom(c echo.Context) error {
	id := c.Param("id")
	idRoom, _ := strconv.Atoi(id)
	data, err := h.feedbackBusiness.GetFeedbackByRoom(idRoom)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", response.FromCoreList(data)))
}
