package business

import "be9/mnroom/features/feedback"

type feedbackUseCase struct {
	feedbackData feedback.Data
}

func NewFeedbackBusiness(fdbk feedback.Data) feedback.Business {
	return &feedbackUseCase{
		feedbackData: fdbk,
	}
}

func (uc *feedbackUseCase) InsertFeedback(insert feedback.Core) (row int, err error) {
	row, err = uc.feedbackData.InsertFeedback(insert)
	return row, err
}

func (uc *feedbackUseCase) GetDataRoom(id int) (data int, err error) {
	data, err = uc.feedbackData.GetDataRoom(id)
	return data, err
}

func (uc *feedbackUseCase) GetDataRent(id int) (data int, err error) {
	data, err = uc.feedbackData.GetDataRent(id)
	return data, err
}

func (uc *feedbackUseCase) GetFeedbackByRoom(id int) (data []feedback.Core, err error) {
	data, err = uc.feedbackData.GetFeedbackByRoom(id)
	return data, err
}
