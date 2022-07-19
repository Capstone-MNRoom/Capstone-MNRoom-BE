package presentation

import (
	"be9/mnroom/features/users"
	"be9/mnroom/features/users/presentation/request"
	"be9/mnroom/features/users/presentation/response"
	"be9/mnroom/helper"
	_middlewares "be9/mnroom/middlewares"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errValidator.Error()))
	}
	newUser := request.ToCore(insertData)
	row, err := h.userBusiness.InsertData(newUser)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("email or telephone number already exist"))
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

func (h *UserHandler) GetData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	data, err := h.userBusiness.GetData(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", response.FromCore(data)))
}

func (h *UserHandler) DeleteData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	row, err := h.userBusiness.DeleteData(idToken)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to deleted data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to deleted data"))
}

func (h *UserHandler) UpdateData(c echo.Context) error {
	link, report, err := helper.AddImageUser(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	data, _ := h.userBusiness.GetData(idToken)
	updatedData := request.User{
		Username: c.FormValue("username"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Phone:    c.FormValue("phone"),
		Address:  c.FormValue("address"),
		Image:    link,
	}
	if updatedData.Image == "https://storage.googleapis.com/event2022/profile_default.png" {
		updatedData.Image = data.Image
	}
	if updatedData.Username == "" {
		updatedData.Username = data.Username
	}
	if updatedData.Email == "" {
		updatedData.Email = data.Email
	}
	if updatedData.Password == "" {
		updatedData.Password = data.Password
	} else if updatedData.Password != "" {
		hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if errHash != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to generate password"))
		}
		updatedData.Password = string(hashedPassword)
	}
	if updatedData.Phone == "" {
		updatedData.Phone = data.Phone
	}
	if updatedData.Address == "" {
		updatedData.Address = data.Address
	}
	v := validator.New()
	errValidator := v.Struct(updatedData)
	if errValidator != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errValidator.Error()))
	}
	newUser := request.ToCore(updatedData)
	row, err := h.userBusiness.UpdateData(idToken, newUser)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to updated data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to updated data"))
}
