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

	// Login
	e.POST("/login", presenter.AuthPresenter.Auth)

	// Categorys
	e.POST("/categorys", presenter.CategoryPresenter.InsertData)
	e.GET("/categorys", presenter.CategoryPresenter.GetDataAll)
	e.GET("/categorys/:id", presenter.CategoryPresenter.GetData)
	e.PUT("/categorys/:id", presenter.CategoryPresenter.UpdateData)
	e.DELETE("/categorys/:id", presenter.CategoryPresenter.DeleteData)

	// Facilitys
	e.POST("/facilitys", presenter.FacilityPresenter.InsertData)
	e.GET("/facilitys", presenter.FacilityPresenter.GetDataAll)
	e.GET("/facilitys/:id", presenter.FacilityPresenter.GetData)
	e.PUT("/facilitys/:id", presenter.FacilityPresenter.UpdateData)
	e.DELETE("/facilitys/:id", presenter.FacilityPresenter.DeleteData)
	
	// Signup
	e.POST("/signup", presenter.UserPresenter.InsertData)
	// Users
	e.GET("/users", presenter.UserPresenter.GetAllData, _middlewares.JWTMiddleware())
	e.GET("/users/profile", presenter.UserPresenter.GetData, _middlewares.JWTMiddleware())
	e.DELETE("/users", presenter.UserPresenter.DeleteData, _middlewares.JWTMiddleware())
	e.PUT("/users", presenter.UserPresenter.UpdateData, _middlewares.JWTMiddleware())
	// e.PUT("/users", presenter.UserPresenter.UpdateData, _middlewares.JWTMiddleware())

	return e
}
