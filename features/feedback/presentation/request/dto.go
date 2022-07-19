package request

import (
	"be9/mnroom/features/feedback"
)

type Feedback struct {
	Rating  int    `json:"rating" validate:"required" form:"rating"`
	Comment string `json:"comment" validate:"required" form:"comment"`
	RentsID uint   `json:"rents_id" validate:"required" form:"rents_id"`
}

func ToCore(req Feedback) feedback.Core {
	return feedback.Core{
		Rating:  req.Rating,
		Comment: req.Comment,
		Rents: feedback.Rents{
			ID: int(req.RentsID),
		},
	}
}
