package response

import (
	"be9/mnroom/features/facilitys"
	"time"
)

type Facilitys struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func FromCoreList(data []facilitys.Core) []Facilitys {
	result := []Facilitys{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCore(data facilitys.Core) Facilitys {
	return Facilitys{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
	}
}
