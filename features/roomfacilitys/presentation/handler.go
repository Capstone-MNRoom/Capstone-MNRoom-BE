package presentation

import (
	"be9/mnroom/features/roomfacilitys"
	"be9/mnroom/features/roomfacilitys/presentation/response"
	"be9/mnroom/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomFacilityHandler struct {
	roomfacilitysBusiness roomfacilitys.Business
}

func NewRoomFacilitysHandler(business roomfacilitys.Business) *RoomFacilityHandler {
	return &RoomFacilityHandler{
		roomfacilitysBusiness: business,
	}
}

func (m *RoomFacilityHandler) GetData(c echo.Context) error {
	id := c.Param("id")
	idRoom, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid id"))
	}
	row, err := m.roomfacilitysBusiness.GetDataRow(idRoom)
	if row != 1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("invalid input"))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get all data"))
	}
	data, _ := m.roomfacilitysBusiness.GetData(idRoom)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success to get all data", response.FromCoreList(data)))
}
