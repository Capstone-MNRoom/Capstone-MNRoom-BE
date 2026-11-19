package business

import (
	"be9/mnroom/features/feedback"
	"context"
)

type feedbackUseCase struct {
	feedbackData feedback.Data
}

func NewFeedbackBusiness(fdbk feedback.Data) feedback.Business {
	return &feedbackUseCase{
		feedbackData: fdbk,
	}
}

func (uc *feedbackUseCase) InsertFeedback(ctx context.Context, insert feedback.Core) (row int, err error) {
	row, err = uc.feedbackData.InsertFeedback(ctx, insert)
	return row, err
}

func (uc *feedbackUseCase) GetDataRoom(ctx context.Context, id int) (data int, err error) {
	data, err = uc.feedbackData.GetDataRoom(ctx, id)
	return data, err
}

func (uc *feedbackUseCase) GetDataRent(ctx context.Context, id int) (data int, err error) {
	data, err = uc.feedbackData.GetDataRent(ctx, id)
	return data, err
}

func (uc *feedbackUseCase) GetFeedbackByRoom(ctx context.Context, id int) (data []feedback.Core, err error) {
	data, err = uc.feedbackData.GetFeedbackByRoom(ctx, id)
	return data, err
}

func (uc *feedbackUseCase) GetDataRentUser(ctx context.Context, idToken int, idRoom int) (data int, err error) {
	data, err = uc.feedbackData.GetDataRentUser(ctx, idToken, idRoom)
	return data, err
}
