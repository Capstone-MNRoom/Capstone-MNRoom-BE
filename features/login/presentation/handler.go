package presentation

import (
	"be9/mnroom/features/login"
	"be9/mnroom/features/login/presentation/request"
	"be9/mnroom/features/login/presentation/response"
	"be9/mnroom/helper"
	_middlewares "be9/mnroom/middlewares"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authBusiness login.Business
}

func NewAuthHandler(business login.Business) *AuthHandler {
	return &AuthHandler{
		authBusiness: business,
	}
}

func (a *AuthHandler) Auth(c echo.Context) error {
	insertLogin := request.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	errBind := c.Bind(&insertLogin)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	v := validator.New()
	errValidator := v.Struct(insertLogin)
	if errValidator != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errValidator.Error()))
	}
	authUser := request.ToCore(insertLogin)
	dataAuth, err := a.authBusiness.Auth(authUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("email or password incorrect"))
	}
	token, errToken := _middlewares.CreateToken(int(dataAuth.ID))
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("error toker"))
	}
	return c.JSON(http.StatusOK,
		map[string]interface{}{
			"message": "success",
			"data":    response.FromCore(dataAuth),
			"token":   token,
		},
	)

}
