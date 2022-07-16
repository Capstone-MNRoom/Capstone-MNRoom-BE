package data

import (
	"be9/mnroom/features/categorys"

	"gorm.io/gorm"
)

type Categorys struct {
	gorm.Model
	CategoryName string `json:"category_name" form:"category_name"`
}

func toCoreList(data []Categorys) []categorys.Core {
	result := []categorys.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *Categorys) toCore() categorys.Core {
	return categorys.Core{
		ID:           int(data.ID),
		CategoryName: data.CategoryName,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func fromCore(core categorys.Core) Categorys {
	return Categorys{
		CategoryName: core.CategoryName,
	}
}
