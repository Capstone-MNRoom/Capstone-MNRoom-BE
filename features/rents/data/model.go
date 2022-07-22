package data

import (
	"be9/mnroom/features/rents"
	"be9/mnroom/features/rooms"
	_rooms "be9/mnroom/features/rooms/data"
	"be9/mnroom/features/users"
	_users "be9/mnroom/features/users/data"

	"gorm.io/gorm"
)

type Rents struct {
	gorm.Model
	DateStart        string       `json:"date_start" form:"date_start"`
	DateEnd          string       `json:"date_end" form:"date_end"`
	Bank             string       `json:"bank" form:"bank"`
	TotalRentalPrice int          `json:"total_rental_price" form:"total_rental_price"`
	Status           string       `json:"status" form:"status"`
	UserID           uint         `json:"user_id" form:"user_id"`
	RoomsID          uint         `json:"rooms_id" form:"rooms_id"`
	User             _users.User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rooms            _rooms.Rooms `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Payments struct {
	gorm.Model
	TransactionID     string `json:"transaction_id" form:"transaction_id"`
	PaymentType       string `json:"payment_type" form:"payment_type"`
	OrderID           int    `json:"order_id" form:"order_id"`
	BankTransfer      string `json:"bank_transfer" form:"bank_transfer"`
	GrossAmount       int    `json:"gross_amount" form:"gross_amount"`
	VANumber          int    `json:"va_number" form:"va_number"`
	TransactionStatus string `json:"transaction_status" form:"transaction_status"`
	UserID            int    `json:"user_id" form:"user_id"`
	RentsID           int    `json:"rents_id" form:"rents_id"`
	User              _users.User
	Rents             Rents
}

func toCoreList(data []Rents) []rents.Core {
	result := []rents.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *Rents) toCore() rents.Core {
	return rents.Core{
		ID:               int(data.ID),
		DateStart:        data.DateStart,
		DateEnd:          data.DateEnd,
		Bank:             data.Bank,
		TotalRentalPrice: data.TotalRentalPrice,
		Status:           data.Status,
		User: users.Core{
			ID:       int(data.User.ID),
			Username: data.User.Username,
			Email:    data.User.Email,
		},
		Room: rooms.Core{
			ID:        int(data.Rooms.ID),
			RoomName:  data.Rooms.RoomName,
			HotelName: data.Rooms.HotelName,
			ImageRoom: data.Rooms.ImageRoom,
			Capacity:  data.Rooms.Capacity,
		},
	}
}

func fromCore(core rents.Core) Rents {
	return Rents{
		DateStart:        core.DateStart,
		DateEnd:          core.DateEnd,
		Bank:             core.Bank,
		TotalRentalPrice: core.TotalRentalPrice,
		Status:           core.Status,
		UserID:           uint(core.User.ID),
		RoomsID:          uint(core.Room.ID),
	}
}

func toCoreListPayments(data []Payments) []rents.CorePayments {
	result := []rents.CorePayments{}
	for key := range data {
		result = append(result, data[key].toCorePayment())
	}
	return result
}

func (data *Payments) toCorePayment() rents.CorePayments {
	return rents.CorePayments{
		ID:                int(data.ID),
		TransactionID:     data.TransactionID,
		PaymentType:       data.PaymentType,
		OrderID:           data.OrderID,
		BankTransfer:      data.BankTransfer,
		GrossAmount:       data.GrossAmount,
		VANumber:          data.VANumber,
		TransactionStatus: data.TransactionStatus,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
		User: users.Core{
			ID:       int(data.User.ID),
			Username: data.User.Username,
			Email:    data.User.Email,
		},
		Rents: rents.Core{
			ID:               int(data.Rents.ID),
			DateStart:        data.Rents.DateStart,
			DateEnd:          data.Rents.DateEnd,
			TotalRentalPrice: data.Rents.TotalRentalPrice,
		},
	}
}

func fromCorePayment(core rents.CorePayments) Payments {
	return Payments{
		TransactionID:     core.TransactionID,
		PaymentType:       core.PaymentType,
		OrderID:           core.OrderID,
		BankTransfer:      core.BankTransfer,
		GrossAmount:       core.GrossAmount,
		VANumber:          core.VANumber,
		TransactionStatus: core.TransactionStatus,
		UserID:            core.User.ID,
		RentsID:           core.Rents.ID,
	}
}
