package request

import (
	"be9/mnroom/features/rents"
	"be9/mnroom/features/rooms"
)

type Rents struct {
	Date_start         string `json:"date_start" validate:"required" form:"date_start"`
	Date_end           string `json:"date_end" validate:"required" form:"date_end"`
	Bank               string `json:"bank" validate:"required" form:"bank"`
	Total_rental_price int    `json:"total_rental_price" validate:"required" form:"total_rental_price"`
	Status             string `json:"status" validate:"required" form:"status"`
	RoomID             uint   `json:"rooms_id" validate:"required" form:"rooms_id"`
}

func ToCore(req Rents) rents.Core {
	return rents.Core{
		Date_start:         req.Date_start,
		Date_end:           req.Date_end,
		Bank:               req.Bank,
		Total_rental_price: req.Total_rental_price,
		Status:             req.Status,
		Room: rooms.Core{
			ID: int(req.RoomID),
		},
	}
}
