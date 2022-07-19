package response

import (
	"be9/mnroom/features/feedback"
	"time"
)

type Feedback struct {
	ID        int       `json:"id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	User      User
	Rents     Rents
}

type User struct {
	ID       int
	Username string
}

type Rents struct {
	ID                 int
	Date_start         string
	Date_end           string
	Bank               string
	Total_rental_price int
	Status             string
	Room               Rooms
}

type Rooms struct {
	ID int
}

func FromCore(data feedback.Core) Feedback {
	return Feedback{
		ID:        data.ID,
		Rating:    data.Rating,
		Comment:   data.Comment,
		CreatedAt: data.CreatedAt,
		User: User{
			ID:       data.User.ID,
			Username: data.User.Username,
		},
		Rents: Rents{
			ID:                 data.Rents.ID,
			Date_start:         data.Rents.Date_start,
			Date_end:           data.Rents.Date_end,
			Bank:               data.Rents.Bank,
			Total_rental_price: data.Rents.Total_rental_price,
			Status:             data.Rents.Status,
			Room:               Rooms(data.Rents.Room),
		},
	}
}

func FromCoreList(data []feedback.Core) []Feedback {
	result := []Feedback{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
