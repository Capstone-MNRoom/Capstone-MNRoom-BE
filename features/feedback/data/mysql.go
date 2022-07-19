package data

import (
	"be9/mnroom/features/feedback"

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

func (repo *mysqlFeedbackRepository) InsertFeedback(insert feedback.Core) (row int, err error) {
	insertFeedback := fromCore(insert)
	tx := repo.db.Create(&insertFeedback)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlFeedbackRepository) GetFeedbackByRoom(id int) (data []feedback.Core, err error) {
	var getAllData []Feedback
	tx := repo.db.Table("Feedback").Where("r.rooms_id = ?", id).Joins("inner join Users u on Feedback.user_id = u.id").Joins("inner join Rents r on f.rents_id = r.id").Find(&getAllData)
	// tx := repo.db.Preload("User").Preload("Rents").Find(&getAllData).Table("Rents").Select("rooms_id(id,?)", id)
	// tx := repo.db.Preload("User").Preload("Rooms").Find(&getAllData)
	if tx.Error != nil {
		return []feedback.Core{}, tx.Error
	}
	return toCoreList(getAllData), nil
}
