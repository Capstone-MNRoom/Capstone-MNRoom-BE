package presentation

import (
	"be9/mnroom/features/rooms"
	"be9/mnroom/features/rooms/presentation/request"
	"be9/mnroom/features/rooms/presentation/response"
	"be9/mnroom/helper"
	_middlewares "be9/mnroom/middlewares"
	"strconv"

	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	roomBusiness rooms.Business
}

func NewRoomHandler(business rooms.Business) *RoomHandler {
	return &RoomHandler{
		roomBusiness: business,
	}
}

func (v *RoomHandler) InsertData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	link, report, err := helper.AddImageRoom(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	linkLogo, reportLogo, err2 := helper.AddImageLogo(c)
	if err2 != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", reportLogo["message"])))
	}
	var insertData request.Rooms
	errBind := c.Bind(&insertData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	insertData.ImageRoom = link
	insertData.ImageLogo = linkLogo
	val := validator.New()
	errValidator := val.Struct(insertData)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	newRoom := request.ToCore(insertData)

	newRoom.User.ID = idToken
	row, err := v.roomBusiness.InsertData(newRoom)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert data"))
}

func (v *RoomHandler) GetAllData(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	dat, err := v.roomBusiness.GetAllData(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	// return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success get all data"))
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(dat)))
}

func (v *RoomHandler) GetData(c echo.Context) error {
	id := c.Param("id")
	idRoom, _ := strconv.Atoi(id)
	dat, err := v.roomBusiness.GetData(idRoom)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get data"))
	}
	// return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to get data"))
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCore(dat)))
}

func (v *RoomHandler) DeleteData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	id := c.Param("id")
	idRoom, _ := strconv.Atoi(id)
	data, _ := v.roomBusiness.GetToken(idRoom, idToken)
	if data.User.ID != idToken {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
	}
	row, err := v.roomBusiness.DeleteData(idRoom)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to deleted data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to deleted data"))
}

func (v *RoomHandler) UpdateData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	id := c.Param("id")
	idEvent, _ := strconv.Atoi(id)
	dataToken, _ := v.roomBusiness.GetToken(idEvent, idToken)
	if dataToken.User.ID != idToken {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
	}
	link, report, err := helper.AddImageRoom(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	linkLogo, reportLogo, err := helper.AddImageLogo(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", reportLogo["message"])))
	}
	data, _ := v.roomBusiness.GetData(idEvent)
	CapacityInt, _ := strconv.Atoi("capacity")
	RentalPriceInt, _ := strconv.Atoi("rental_price")
	categorysIDInt, _ := strconv.Atoi(c.FormValue("categorys_id"))
	updatedData := request.Rooms{
		CategorysID: uint(categorysIDInt),
		ImageRoom:   link,
		ImageLogo:   linkLogo,
		RoomName:    c.FormValue("room_name"),
		Capacity:    CapacityInt,
		RentalPrice: RentalPriceInt,
		City:        c.FormValue("city"),
		Address:     c.FormValue("address"),
	}
	if updatedData.CategorysID == 0 {
		updatedData.CategorysID = uint(data.Categorys.ID)
	}
	if updatedData.ImageRoom == "https://storage.googleapis.com/event2022/event-gomeet.png" {
		updatedData.ImageRoom = data.ImageRoom
	}
	if updatedData.ImageLogo == "https://storage.googleapis.com/event2022/event-gomeet.png" {
		updatedData.ImageLogo = data.ImageLogo
	}
	if updatedData.RoomName == "" {
		updatedData.RoomName = data.RoomName
	}
	if updatedData.Address == "" {
		updatedData.Address = data.Address
	}
	if updatedData.Capacity == 0 {
		updatedData.Capacity = data.Capacity
	}
	if updatedData.RentalPrice == 0 {
		updatedData.RentalPrice = data.RentalPrice
	}
	if updatedData.City == "" {
		updatedData.City = data.City
	}

	val := validator.New()
	errValidator := val.Struct(updatedData)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	newEvent := request.ToCore(updatedData)
	row, err := v.roomBusiness.UpdatedData(idEvent, newEvent)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to updated data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to updated data"))
}
