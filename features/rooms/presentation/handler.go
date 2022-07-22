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
	errRoomName := v.Var(insertRoom.RoomName, "required")
	if errRoomName != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("room name cannot be empty"))
	}
	errCapacity := v.Var(insertRoom.Capacity, "required,numeric")
	if errCapacity != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("capacity must be a number"))
	}
	errHotelName := v.Var(insertRoom.HotelName, "required")
	if errHotelName != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("hotel name cannot be empty"))
	}
	errRentalPrice := v.Var(insertRoom.RentalPrice, "required,numeric")
	if errRentalPrice != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("rental price must be a number"))
	}
	errAddress := v.Var(insertRoom.Address, "required")
	if errAddress != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("address cannot be empty"))
	}
	errCity := v.Var(insertRoom.City, "required")
	if errCity != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("city cannot be empty"))
	}
	errCategorysID := v.Var(insertRoom.CategorysID, "required,numeric")
	if errCategorysID != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("category cannot be empty"))
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
	idRoom, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid id"))
	}
	data, err := r.roomBusiness.GetData(idRoom)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid input"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get data", response.FromCore(data)))
}

func (r *RoomHandler) UpdateData(c echo.Context) error {
	id := c.Param("id")
	idRoom, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid id"))
	}
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
	v := validator.New()
	if updateData.CategorysID == 0 {
		updateData.CategorysID = uint(data.Categorys.ID)
	} else if updateData.CategorysID != 0 {
		errCategorysID := v.Var(updateData.CategorysID, "required,numeric")
		if errCategorysID != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("category cannot be empty"))
		}
	}
	if updateData.ImageRoom == "https://storage.googleapis.com/event2022/event-gomeet.png" {
		updateData.ImageRoom = data.ImageRoom
	}
	if updateData.ImagePengelola == "https://storage.googleapis.com/event2022/event-gomeet.png" {
		updateData.ImagePengelola = data.ImagePengelola
	}
	if updateData.RoomName == "" {
		updateData.RoomName = data.RoomName
	} else if updateData.RoomName != "" {
		errRoomName := v.Var(updateData.RoomName, "required")
		if errRoomName != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("room name cannot be empty"))
		}
	}
	if updateData.Capacity == 0 {
		updateData.Capacity = data.Capacity
	} else if updateData.Capacity != 0 {
		errCapacity := v.Var(updateData.Capacity, "required,numeric")
		if errCapacity != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("capacity must be a number"))
		}
	}
	if updateData.HotelName == "" {
		updateData.HotelName = data.HotelName
	} else if updateData.HotelName != "" {
		errHotelName := v.Var(updateData.HotelName, "required")
		if errHotelName != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("hotel name cannot be empty"))
		}
	}
	if updateData.RentalPrice == 0 {
		updateData.RentalPrice = data.RentalPrice
	} else if updateData.RentalPrice != 0 {
		errRentalPrice := v.Var(updateData.RentalPrice, "required,numeric")
		if errRentalPrice != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("rental price must be a number"))
		}
	}
	if updateData.Address == "" {
		updateData.Address = data.Address
	} else if updateData.Address != "" {
		errAddress := v.Var(updateData.Address, "required")
		if errAddress != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("address cannot be empty"))
		}
	}
	if updateData.City == "" {
		updateData.City = data.City
	} else if updateData.City != "" {
		errCity := v.Var(updateData.City, "required")
		if errCity != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("city cannot be empty"))
		}
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
	idRoom, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid id"))
	}
	idToken, errToken := _middlewares.ExtractToken(c)
	if errToken != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid token"))
	}
	rowIDRoom, _ := r.roomBusiness.GetDataIDRoom(idRoom)
	dataToken, _ := r.roomBusiness.GetToken(idRoom, idToken)
	if rowIDRoom != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid input"))
	} else if rowIDRoom == 1 && dataToken.User.ID != idToken {
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

func (r *RoomHandler) GetDataByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	categoryInt, _ := strconv.Atoi(category)
	data, err := r.roomBusiness.GetDataByCategory(categoryInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Check your input"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success to get data", data))
}
