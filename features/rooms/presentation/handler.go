package presentation

import (
	"be9/mnroom/features/rooms"
	"be9/mnroom/features/rooms/presentation/request"
	"be9/mnroom/features/rooms/presentation/response"
	"be9/mnroom/helper"
	_middlewares "be9/mnroom/middlewares"
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
	var insertRoom request.Rooms
	errBind := c.Bind(&insertRoom)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data, check your input"))
	}
	insertRoom.ImageRoom = "gambar1.jpg"
	insertRoom.ImagePengelola = "gambar2.jpg"
	v := validator.New()
	errValidator := v.Struct(insertRoom)
	if errValidator != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(errValidator.Error()))
	}
	newRoom := request.ToCore(insertRoom)
	newRoom.User.ID = idToken
	row, err := r.roomBusiness.InsertData(newRoom)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to insert data"))
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
