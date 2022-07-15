package presentation

import (
	"be9/mnroom/features/users"
	"be9/mnroom/features/users/presentation/response"
	"be9/mnroom/features/users/presentation/request"
	"be9/mnroom/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: business,
	}
}

func (h *UserHandler) InsertData(c echo.Context) error {
	link, report, err := helper.AddImageUser(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	var insertData request.User
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	insertData.Image = link
	v := validator.New()
	errValidator := v.Struct(insertData)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	newUser := request.ToCore(insertData)
	row, err := h.userBusiness.InsertData(newUser)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert data"))
}
func (h *UserHandler) GetAllData(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	data, err := h.userBusiness.GetAllData(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}
