package data

import (
	"be9/mnroom/features/facilitys"

	"gorm.io/gorm"
)

type mysqlFacilityRepository struct {
	db *gorm.DB
}

func NewFacilityRepository(conn *gorm.DB) facilitys.Data {
	return &mysqlFacilityRepository{
		db: conn,
	}
}

func (repo *mysqlFacilityRepository) InsertData(insert facilitys.Core) (row int, err error) {
	insertFacility := fromCore(insert)
	tx := repo.db.Create(&insertFacility)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlFacilityRepository) GetDataAll() (data []facilitys.Core, err error) {
	var getDataAll []Facilitys
	tx := repo.db.Find(&getDataAll)
	if tx.Error != nil {
		return []facilitys.Core{}, tx.Error
	}
	return toCoreList(getDataAll), nil
}

func (repo *mysqlFacilityRepository) GetData(id int) (data facilitys.Core, err error) {
	var getData Facilitys
	tx := repo.db.First(&getData, id)
	if tx.Error != nil {
		return facilitys.Core{}, tx.Error
	}
	return getData.toCore(), nil
}

func (repo *mysqlFacilityRepository) UpdateData(id int, insert facilitys.Core) (row int, err error) {
	tx := repo.db.Model(&Facilitys{}).Where("id = ?", id).Updates(Facilitys{Name: insert.Name})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlFacilityRepository) DeleteData(id int) (row int, err error) {
	var deleteData Facilitys
	tx := repo.db.Unscoped().Delete(&deleteData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
