package data

import (
	"be9/mnroom/features/rooms"
	"context"

	"gorm.io/gorm"
)

type mysqlRoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(conn *gorm.DB) rooms.Data {
	return &mysqlRoomRepository{
		db: conn,
	}
}

func (repo *mysqlRoomRepository) InsertData(ctx context.Context, insert rooms.Core) (data rooms.Core, err error) {
	insertRoom := fromCore(insert)
	tx := repo.db.WithContext(ctx).Create(&insertRoom)
	if tx.Error != nil {
		return rooms.Core{}, tx.Error
	}
	return insertRoom.toCore(), nil
}

func (repo *mysqlRoomRepository) InsertDataRoomFacilitys(ctx context.Context, insert rooms.CoreRoomFacilitys) (row int, err error) {
	tx := repo.db.WithContext(ctx).Model(&RoomFacilitys{}).Create([]map[string]interface{}{
		{"user_id": insert.User.ID, "rooms_id": insert.Rooms.ID, "facilitys_id": insert.Facilitys.ID},
	})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRoomRepository) GetDataAll(ctx context.Context, page int) (data []rooms.Core, err error) {
	var getDataAll []Rooms

	tx := repo.db.WithContext(ctx).Limit(12).Offset(page).Preload("User").Preload("Categorys").Find(&getDataAll)
	if tx.Error != nil {
		return []rooms.Core{}, tx.Error
	}
	return toCoreList(getDataAll), nil
}

func (repo *mysqlRoomRepository) GetData(ctx context.Context, id int) (data rooms.Core, err error) {
	var getData Rooms
	tx := repo.db.WithContext(ctx).Preload("User").Preload("Categorys").First(&getData, id)
	if tx.Error != nil {
		return rooms.Core{}, tx.Error
	}
	return getData.toCore(), nil
}

func (repo *mysqlRoomRepository) GetDataIDRoom(ctx context.Context, id int) (row int, err error) {
	var getData Rooms
	tx := repo.db.WithContext(ctx).Preload("User").Preload("Categorys").First(&getData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRoomRepository) UpdateData(ctx context.Context, id int, insert rooms.Core) (row int, err error) {
	tx := repo.db.WithContext(ctx).Model(&Rooms{}).Where("id = ?", id).Updates(Rooms{ImageRoom: insert.ImageRoom, ImagePengelola: insert.ImagePengelola, RoomName: insert.RoomName, Capacity: insert.Capacity, HotelName: insert.HotelName, RentalPrice: insert.RentalPrice, Address: insert.Address, City: insert.City, CategorysID: uint(insert.Categorys.ID)})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRoomRepository) DeleteData(ctx context.Context, id int) (row int, err error) {
	var deleteData Rooms
	tx := repo.db.WithContext(ctx).Unscoped().Delete(&deleteData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRoomRepository) GetToken(ctx context.Context, id int, idToken int) (data rooms.Core, err error) {
	var getData Rooms
	tx := repo.db.WithContext(ctx).Preload("User").Preload("Categorys").First(&getData, id)
	if tx.Error != nil {
		return rooms.Core{}, tx.Error
	}
	if getData.toCore().User.ID != idToken {
		return rooms.Core{}, err
	}
	return getData.toCore(), nil
}

func (repo *mysqlRoomRepository) GetDataAllUserRoom(ctx context.Context, idToken int) (data []rooms.Core, err error) {
	var getDataUserRoom []Rooms
	tx := repo.db.WithContext(ctx).Where("user_id = ?", idToken).Preload("User").Preload("Categorys").Find(&getDataUserRoom)
	if tx.Error != nil {
		return []rooms.Core{}, tx.Error
	}
	return toCoreList(getDataUserRoom), nil
}

func (repo *mysqlRoomRepository) GetDataByCategory(ctx context.Context, id int) (data []rooms.Core, err error) {
	var categoryId []Rooms
	tx := repo.db.WithContext(ctx).Where("categorys_id = ? ", id).Preload("User").Preload("Categorys").Find(&categoryId)
	if tx.Error != nil {
		return []rooms.Core{}, tx.Error
	}
	return toCoreList(categoryId), nil
}
