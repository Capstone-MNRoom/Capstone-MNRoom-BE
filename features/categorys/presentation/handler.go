package presentation

import (
	"be9/mnroom/features/categorys"
	"be9/mnroom/features/categorys/presentation/request"
	"be9/mnroom/features/categorys/presentation/response"
	"be9/mnroom/helper"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryBusiness categorys.Business
}

func NewCategoryHandler(business categorys.Business) *CategoryHandler {
	return &CategoryHandler{
		categoryBusiness: business,
	}
}

func (g *CategoryHandler) InsertData(c echo.Context) error {
	var insertCategory request.Categorys
	errBind := c.Bind(&insertCategory)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	v := validator.New()
	errValidator := v.Struct(insertCategory)
	if errValidator != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errValidator.Error()))
	}
	newCategory := request.ToCore(insertCategory)
	row, err := g.categoryBusiness.InsertData(newCategory)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert data"))
}

func (g *CategoryHandler) GetDataAll(c echo.Context) error {
	data, err := g.categoryBusiness.GetDataAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}

func (g *CategoryHandler) GetData(c echo.Context) error {
	id := c.Param("id")
	idCategory, _ := strconv.Atoi(id)
	data, err := g.categoryBusiness.GetData(idCategory)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", response.FromCore(data)))
}

func (g *CategoryHandler) UpdateData(c echo.Context) error {
	id := c.Param("id")
	idCategory, _ := strconv.Atoi(id)
	updateCategory := request.Categorys{
		CategoryName: c.FormValue("category_name"),
	}
	v := validator.New()
	errValidator := v.Struct(updateCategory)
	if errValidator != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errValidator.Error()))
	}
	newUpdate := request.ToCore(updateCategory)
	row, err := g.categoryBusiness.UpdateData(idCategory, newUpdate)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to updated data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to updated data"))
}

func (g *CategoryHandler) DeleteData(c echo.Context) error {
	id := c.Param("id")
	idCategory, _ := strconv.Atoi(id)
	row, err := g.categoryBusiness.DeleteData(idCategory)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to deleted data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to deleted data"))
}
