package request

import "be9/mnroom/features/facilitys"

type Facilitys struct {
	Name string `json:"name" validate:"required" form:"name"`
}

func ToCore(req Facilitys) facilitys.Core {
	return facilitys.Core{
		Name: req.Name,
	}
}
