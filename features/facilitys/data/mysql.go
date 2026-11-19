package data

import (
	"be9/mnroom/features/facilitys"
	"context"

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

func (repo *mysqlFacilityRepository) InsertData(ctx context.Context, insert facilitys.Core) (row int, err error) {
	insertFacility := fromCore(insert)
	tx := repo.db.WithContext(ctx).Create(&insertFacility)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlFacilityRepository) GetDataAll(ctx context.Context) (data []facilitys.Core, err error) {
	var getDataAll []Facilitys
	tx := repo.db.WithContext(ctx).Find(&getDataAll)
	if tx.Error != nil {
		return []facilitys.Core{}, tx.Error
	}
	return toCoreList(getDataAll), nil
}

func (repo *mysqlFacilityRepository) GetData(ctx context.Context, id int) (data facilitys.Core, err error) {
	var getData Facilitys
	tx := repo.db.WithContext(ctx).First(&getData, id)
	if tx.Error != nil {
		return facilitys.Core{}, tx.Error
	}
	return getData.toCore(), nil
}

func (repo *mysqlFacilityRepository) UpdateData(ctx context.Context, id int, insert facilitys.Core) (row int, err error) {
	tx := repo.db.WithContext(ctx).Model(&Facilitys{}).Where("id = ?", id).Updates(Facilitys{Name: insert.Name})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlFacilityRepository) DeleteData(ctx context.Context, id int) (row int, err error) {
	var deleteData Facilitys
	tx := repo.db.WithContext(ctx).Unscoped().Delete(&deleteData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
