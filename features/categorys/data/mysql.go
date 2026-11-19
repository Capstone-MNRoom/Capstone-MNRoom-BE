package data

import (
	"be9/mnroom/features/categorys"
	"context"

	"gorm.io/gorm"
)

type mysqlCategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(conn *gorm.DB) categorys.Data {
	return &mysqlCategoryRepository{
		db: conn,
	}
}

func (repo *mysqlCategoryRepository) InsertData(ctx context.Context, insert categorys.Core) (row int, err error) {
	insertCategory := fromCore(insert)
	tx := repo.db.WithContext(ctx).Create(&insertCategory)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlCategoryRepository) GetDataAll(ctx context.Context) (data []categorys.Core, err error) {
	var getDataAll []Categorys
	tx := repo.db.WithContext(ctx).Find(&getDataAll)
	if tx.Error != nil {
		return []categorys.Core{}, tx.Error
	}
	return toCoreList(getDataAll), nil
}

func (repo *mysqlCategoryRepository) GetData(ctx context.Context, id int) (data categorys.Core, err error) {
	var getData Categorys
	tx := repo.db.WithContext(ctx).First(&getData, id)
	if tx.Error != nil {
		return categorys.Core{}, tx.Error
	}
	return getData.toCore(), nil
}

func (repo *mysqlCategoryRepository) UpdateData(ctx context.Context, id int, insert categorys.Core) (row int, err error) {
	tx := repo.db.WithContext(ctx).Model(&Categorys{}).Where("id = ?", id).Updates(Categorys{CategoryName: insert.CategoryName})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlCategoryRepository) DeleteData(ctx context.Context, id int) (row int, err error) {
	var deleteData Categorys
	tx := repo.db.WithContext(ctx).Unscoped().Delete(&deleteData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
