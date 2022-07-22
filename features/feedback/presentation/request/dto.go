package request

import (
	"be9/mnroom/features/feedback"
)

type Feedback struct {
	Rating  int    `json:"rating" form:"rating"`
	Comment string `json:"comment" form:"comment"`
	RentsID uint   `json:"rents_id" form:"rents_id"`
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
