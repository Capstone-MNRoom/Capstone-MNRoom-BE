package data

import (
	"be9/mnroom/features/feedback"
	_rents "be9/mnroom/features/rents/data"
	_rooms "be9/mnroom/features/rooms/data"
	"context"

	"gorm.io/gorm"
)

type mysqlFeedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(conn *gorm.DB) feedback.Data {
	return &mysqlFeedbackRepository{
		db: conn,
	}
}

func (repo *mysqlFeedbackRepository) InsertFeedback(ctx context.Context, insert feedback.Core) (row int, err error) {
	insertFeedback := fromCore(insert)
	tx := repo.db.WithContext(ctx).Create(&insertFeedback)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlFeedbackRepository) GetDataRoom(ctx context.Context, id int) (data int, err error) {
	var getDataRoom _rooms.Rooms
	tx := repo.db.WithContext(ctx).First(&getDataRoom, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(getDataRoom.ID), nil
}

func (repo *mysqlFeedbackRepository) GetDataRent(ctx context.Context, id int) (data int, err error) {
	var getDataRent _rents.Rents
	tx := repo.db.WithContext(ctx).Where("rooms_id = ?", id).Preload("Rooms").First(&getDataRent)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(getDataRent.ID), nil
}

func (repo *mysqlFeedbackRepository) GetFeedbackByRoom(ctx context.Context, id int) (data []feedback.Core, err error) {
	var getFeedbackByRoom []Feedback
	tx := repo.db.WithContext(ctx).Where("rents_id = ?", id).Preload("User").Preload("Rents").Find(&getFeedbackByRoom)
	if tx.Error != nil {
		return []feedback.Core{}, tx.Error
	}
	return toCoreList(getFeedbackByRoom), nil
}

func (repo *mysqlFeedbackRepository) GetDataRentUser(ctx context.Context, idToken int, idRoom int) (data int, err error) {
	var getDataRent _rents.Rents
	tx := repo.db.WithContext(ctx).Where("user_id = ? AND rooms_id = ?", idToken, idRoom).Preload("User").First(&getDataRent)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(getDataRent.ID), nil
}
