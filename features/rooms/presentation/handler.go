package presentation

import (
	_requestFacilitys "be9/mnroom/features/roomfacilitys/presentation/request"
	"be9/mnroom/features/rooms"
	"be9/mnroom/features/rooms/presentation/request"
	"be9/mnroom/features/rooms/presentation/response"
	"be9/mnroom/helper"
	_middlewares "be9/mnroom/middlewares"
	"fmt"
	"net/http"
	"strconv"

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

func (r *RoomHandler) InsertData(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	link, report, err := helper.AddImageRoom(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	linkPengelola, reportPengelola, errPengelola := helper.AddImagePengelola(c)
	if errPengelola != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", reportPengelola["message"])))
	}
	var insertRoom request.Rooms
	errBind := c.Bind(&insertRoom)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	insertRoom.ImageRoom = link
	insertRoom.ImagePengelola = linkPengelola
	v := validator.New()
	errValidator := v.Struct(insertRoom)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	newRoom := request.ToCore(insertRoom)
	newRoom.User.ID = idToken
	data, err := r.roomBusiness.InsertData(newRoom)
	var insertRoomFacilitys _requestFacilitys.RoomFacilitys
	for _, v := range newRoom.Facilitys {
		newRoomFacilitys := _requestFacilitys.ToCore(insertRoomFacilitys)
		newRoomFacilitys.User.ID = idToken
		newRoomFacilitys.Rooms.ID = data.ID
		newRoomFacilitys.Facilitys.ID = v
		row, _ := r.roomBusiness.InsertDataRoomFacilitys(rooms.CoreRoomFacilitys(newRoomFacilitys))
		if row != 1 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert data"))
		}
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to insert data"))
}

func (r *RoomHandler) GetDataAll(c echo.Context) error {
	data, err := r.roomBusiness.GetDataAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}

func (r *RoomHandler) GetData(c echo.Context) error {
	id := c.Param("id")
	idRoom, _ := strconv.Atoi(id)
	data, err := r.roomBusiness.GetData(idRoom)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", response.FromCore(data)))
}

func (r *RoomHandler) UpdateData(c echo.Context) error {
	id := c.Param("id")
	idRoom, _ := strconv.Atoi(id)
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	dataToken, _ := r.roomBusiness.GetToken(idRoom, idToken)
	if dataToken.User.ID != idToken {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
	}
	link, report, err := helper.AddImageRoom(c)
	if err != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", report["message"])))
	}
	linkPengelola, reportPengelola, errPengelola := helper.AddImagePengelola(c)
	if errPengelola != nil {
		return c.JSON(report["code"].(int), helper.ResponseFailed(fmt.Sprintf("%s", reportPengelola["message"])))
	}
	data, _ := r.roomBusiness.GetData(idRoom)
	CapacityInt, _ := strconv.Atoi(c.FormValue("capacity"))
	RentalPriceInt, _ := strconv.Atoi(c.FormValue("rental_price"))
	categorysIDInt, _ := strconv.Atoi(c.FormValue("categorys_id"))
	updateData := request.Rooms{
		ImageRoom:      link,
		ImagePengelola: linkPengelola,
		RoomName:       c.FormValue("room_name"),
		Capacity:       CapacityInt,
		HotelName:      c.FormValue("hotel_name"),
		RentalPrice:    RentalPriceInt,
		Address:        c.FormValue("address"),
		City:           c.FormValue("city"),
		CategorysID:    uint(categorysIDInt),
	}
	if updateData.CategorysID == 0 {
		updateData.CategorysID = uint(data.Categorys.ID)
	}
	if updateData.ImageRoom == "https://storage.googleapis.com/event2022/event-gomeet.png" {
		updateData.ImageRoom = data.ImageRoom
	}
	if updateData.ImagePengelola == "https://storage.googleapis.com/event2022/event-gomeet.png" {
		updateData.ImagePengelola = data.ImagePengelola
	}
	if updateData.RoomName == "" {
		updateData.RoomName = data.RoomName
	}
	if updateData.Capacity == 0 {
		updateData.Capacity = data.Capacity
	}
	if updateData.HotelName == "" {
		updateData.HotelName = data.HotelName
	}
	if updateData.RentalPrice == 0 {
		updateData.RentalPrice = data.RentalPrice
	}
	if updateData.Address == "" {
		updateData.Address = data.Address
	}
	if updateData.City == "" {
		updateData.City = data.City
	}
	val := validator.New()
	errValidator := val.Struct(updateData)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	newRoom := request.ToCore(updateData)
	row, errRoom := r.roomBusiness.UpdateData(idRoom, newRoom)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to updated data"))
	}
	if errRoom != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to updated data"))
}

func (r *RoomHandler) DeleteData(c echo.Context) error {
	id := c.Param("id")
	idRoom, _ := strconv.Atoi(id)
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	dataToken, _ := r.roomBusiness.GetToken(idRoom, idToken)
	if dataToken.User.ID != idToken {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
	}
	row, err := r.roomBusiness.DeleteData(idRoom)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to deleted data"))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("success to deleted data"))
}

func (r *RoomHandler) GetDataAllUserRoom(c echo.Context) error {
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	data, err := r.roomBusiness.GetDataAllUserRoom(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", response.FromCoreList(data)))
}
