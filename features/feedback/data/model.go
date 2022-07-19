package data

import (
	"be9/mnroom/features/feedback"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Rating  int    `json:"rating" form:"rating"`
	Comment string `json:"comment" form:"comment"`
	UserID  uint   `json:"user_id" form:"user_id"`
	RentsID uint   `json:"rents_id" form:"rents_id"`
	User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rents   Rents  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	gorm.Model
	Image    string `json:"image" form:"image"`
	Username string `json:"username" form:"username"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `gorm:"unique" json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Feedback []Feedback
}

type Rents struct {
	gorm.Model
	Date_start         string `json:"date_start" form:"date_start"`
	Date_end           string `json:"date_end" form:"date_end"`
	Bank               string `json:"bank" form:"bank"`
	Total_rental_price int    `json:"total_rental_price" form:"total_rental_price"`
	Status             string `json:"status" form:"status"`
	UserID             uint   `json:"user_id" form:"user_id"`
	RoomsID            uint   `json:"rooms_id" form:"rooms_id"`
	Rooms              Rooms  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Feedback           []Feedback
}

type Rooms struct {
	gorm.Model
	ImageRoom      string `json:"image_room" form:"image_room"`
	ImagePengelola string `json:"image_pengelola" form:"image_pengelola"`
	Name           string `json:"name" form:"name"`
	Capacity       int    `json:"capacity" form:"capacity"`
	RentalPrice    int    `json:"rental_price" form:"rental_price"`
	Address        string `json:"address" form:"address"`
	City           string `json:"city" form:"city"`
	UserID         uint   `json:"user_id" form:"user_id"`
	CategorysID    uint   `json:"categorys_id" form:"categorys_id"`
	Rents          []Rents
}

func (data *Feedback) toCore() feedback.Core {
	return feedback.Core{
		ID:        int(data.ID),
		Rating:    data.Rating,
		Comment:   data.Comment,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		User: feedback.User{
			ID:       int(data.User.ID),
			Username: data.User.Username,
		},
		Rents: feedback.Rents{
			ID:                 int(data.Rents.ID),
			Date_start:         data.Rents.Date_start,
			Date_end:           data.Rents.Date_end,
			Bank:               data.Rents.Bank,
			Total_rental_price: data.Rents.Total_rental_price,
			Status:             data.Rents.Status,
			Room: feedback.Rooms{
				ID: int(data.Rents.RoomsID),
			},
		},
	}
}

func toCoreList(data []Feedback) []feedback.Core {
	result := []feedback.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core feedback.Core) Feedback {
	return Feedback{
		Rating:  core.Rating,
		Comment: core.Comment,
		UserID:  uint(core.User.ID),
		RentsID: uint(core.Rents.ID),
	}
}
