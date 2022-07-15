package routes

import (
	"be9/mnroom/factory"
	_middlewares "be9/mnroom/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))

	e.Pre(middleware.RemoveTrailingSlash())

	// Users
	e.POST("/users", presenter.UserPresenter.InsertData, _middlewares.JWTMiddleware())
	e.GET("/users", presenter.UserPresenter.GetAllData, _middlewares.JWTMiddleware())
	e.GET("/users/profile", presenter.UserPresenter.GetData, _middlewares.JWTMiddleware())
	e.DELETE("/users", presenter.UserPresenter.DeleteData, _middlewares.JWTMiddleware())
	e.PUT("/users", presenter.UserPresenter.UpdateData, _middlewares.JWTMiddleware())
	// e.PUT("/users", presenter.UserPresenter.UpdateData, _middlewares.JWTMiddleware())

	return e
}
