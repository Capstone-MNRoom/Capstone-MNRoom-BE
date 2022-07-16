package response

import (
	"be9/mnroom/features/categorys"
	"time"
)

type Categorys struct {
	ID           int       `json:"id"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
}

func FromCoreList(data []categorys.Core) []Categorys {
	result := []Categorys{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data categorys.Core) Categorys {
	return Categorys{
		ID:           data.ID,
		CategoryName: data.CategoryName,
		CreatedAt:    data.CreatedAt,
	}
}
