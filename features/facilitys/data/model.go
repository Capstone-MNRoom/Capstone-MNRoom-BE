package data

import (
	"be9/mnroom/features/facilitys"

	"gorm.io/gorm"
)

type Facilitys struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}

func toCoreList(data []Facilitys) []facilitys.Core {
	result := []facilitys.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func (data *Facilitys) toCore() facilitys.Core {
	return facilitys.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func fromCore(core facilitys.Core) Facilitys {
	return Facilitys{
		Name: core.Name,
	}
}
