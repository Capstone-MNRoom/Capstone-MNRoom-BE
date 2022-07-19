package request

import (
	"be9/mnroom/features/rents"
	"be9/mnroom/features/rooms"
)

type Rents struct {
	DateStart        string `json:"date_start" validate:"required" form:"date_start"`
	DateEnd          string `json:"date_end" validate:"required" form:"date_end"`
	Bank             string `json:"bank" validate:"required" form:"bank"`
	TotalRentalPrice int    `json:"total_rental_price" form:"total_rental_price"`
	Status           string `json:"status" form:"status"`
	RoomID           uint   `json:"rooms_id" form:"rooms_id"`
}

func ToCore(req Rents) rents.Core {
	return rents.Core{
		DateStart:        req.DateStart,
		DateEnd:          req.DateEnd,
		Bank:             req.Bank,
		TotalRentalPrice: req.TotalRentalPrice,
		Status:           req.Status,
		Room: rooms.Core{
			ID: int(req.RoomID),
		},
	}
}
