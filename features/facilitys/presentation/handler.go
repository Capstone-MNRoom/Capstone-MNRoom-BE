package presentation

import (
	"be9/mnroom/features/facilitys"
	"be9/mnroom/features/facilitys/presentation/request"
	"be9/mnroom/features/facilitys/presentation/response"
	"be9/mnroom/helper"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type FacilityHandler struct {
	facilityBusiness facilitys.Business
}

func NewFacilityHandler(business facilitys.Business) *FacilityHandler {
	return &FacilityHandler{
		facilityBusiness: business,
	}
}

func (f *FacilityHandler) InsertData(c echo.Context) error {
	var insertFacility request.Facilitys
	errBind := c.Bind(&insertFacility)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	v := validator.New()
	errName := v.Var(insertFacility.Name, "required")
	if errName != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("name cannot be empty"))
	}
	newFacility := request.ToCore(insertFacility)
	row, err := f.facilityBusiness.InsertData(newFacility)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert data"))
}

func (f *FacilityHandler) GetDataAll(c echo.Context) error {
	data, err := f.facilityBusiness.GetDataAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}

func (f *FacilityHandler) GetData(c echo.Context) error {
	id := c.Param("id")
	idFacility, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid id"))
	}
	data, err := f.facilityBusiness.GetData(idFacility)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid input"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", response.FromCore(data)))
}

func (f *FacilityHandler) UpdateData(c echo.Context) error {
	id := c.Param("id")
	idFacility, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid id"))
	}
	updateFacility := request.Facilitys{
		Name: c.FormValue("name"),
	}
	v := validator.New()
	errValidator := v.Struct(updateFacility)
	if errValidator != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errValidator.Error()))
	}
	newUpdate := request.ToCore(updateFacility)
	row, err := f.facilityBusiness.UpdateData(idFacility, newUpdate)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to updated data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to updated data"))
}

func (f *FacilityHandler) DeleteData(c echo.Context) error {
	id := c.Param("id")
	idFacility, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid id"))
	}
	row, err := f.facilityBusiness.DeleteData(idFacility)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to deleted data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to deleted data"))
}
