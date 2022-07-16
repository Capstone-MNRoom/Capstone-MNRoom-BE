package request

import "be9/mnroom/features/categorys"

type Categorys struct {
	CategoryName string `json:"category_name" validate:"required" form:"category_name"`
}

func ToCore(req Categorys) categorys.Core {
	return categorys.Core{
		CategoryName: req.CategoryName,
	}
}
