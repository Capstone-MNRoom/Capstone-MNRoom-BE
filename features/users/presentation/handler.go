package presentation

import (
	"be9/mnroom/features/users"
	"be9/mnroom/features/users/presentation/request"
	"be9/mnroom/features/users/presentation/response"
	"be9/mnroom/helper"
	_middlewares "be9/mnroom/middlewares"
	"fmt"
	"net/http"

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
	errUsername := v.Var(insertData.Username, "required,alpha")
	if errUsername != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("username can only contains alphabet"))
	}
	if len(insertData.Username) < 3 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("username must be at least 3 characters"))
	}
	errEmail := v.Var(insertData.Email, "required,email")
	if errEmail != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid format email"))
	}
	if len(insertData.Password) < 6 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("password must be at least 6 characters"))
	}
	if len(insertData.Phone) < 8 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("phone nummber must be at least 8 numbers"))
	}
	errPhone := v.Var(insertData.Phone, "required,numeric")
	if errPhone != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("phone number must be in numeric"))
	}
	errAddress := v.Var(insertData.Address, "required")
	if errAddress != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("address cannot be empty"))
	}
	newUser := request.ToCore(insertData)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashedPassword)
	row, err := h.userBusiness.InsertData(newUser)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("email or phone number already exist"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert data"))
}

func (h *UserHandler) GetAllData(c echo.Context) error {
	data, err := h.userBusiness.GetAllData()
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
		Password: c.FormValue("password"),
		Address:  c.FormValue("address"),
		Image:    link,
	}
	v := validator.New()
	if updatedData.Image == "https://storage.googleapis.com/event2022/profile_default.png" {
		updatedData.Image = data.Image
	}
	if updatedData.Username == "" {
		updatedData.Username = data.Username
	} else if updatedData.Username != "" {
		errUsername := v.Var(updatedData.Username, "required,alpha")
		if errUsername != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("username can only contains alphabet"))
		}
		if len(updatedData.Username) < 3 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("username must be at least 3 characters"))
		}
	}
	if updatedData.Password == "" {
		updatedData.Password = data.Password
	} else if updatedData.Password != "" {
		if len(updatedData.Password) < 6 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("password must be at least 6 characters"))
		}
		hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if errHash != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to generate password"))
		}
		updatedData.Password = string(hashedPassword)
	}
	if updatedData.Address == "" {
		updatedData.Address = data.Address
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
