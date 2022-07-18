package data

import (
	"be9/mnroom/features/rooms"

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

func (repo *mysqlRoomRepository) InsertData(insert rooms.Core) (data rooms.Core, err error) {
	insertRoom := fromCore(insert)
	tx := repo.db.Create(&insertRoom)
	if tx.Error != nil {
		return rooms.Core{}, tx.Error
	}
	return insertRoom.toCore(), nil
}

func (repo *mysqlRoomRepository) InsertDataRoomFacilitys(insert rooms.CoreRoomFacilitys) (row int, err error) {
	tx := repo.db.Model(&RoomFacilitys{}).Create([]map[string]interface{}{
		{"user_id": insert.User.ID, "rooms_id": insert.Rooms.ID, "facilitys_id": insert.Facilitys.ID},
	})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRoomRepository) GetDataAll() (data []rooms.Core, err error) {
	var getDataAll []Rooms
	tx := repo.db.Preload("User").Preload("Categorys").Find(&getDataAll)
	if tx.Error != nil {
		return []rooms.Core{}, tx.Error
	}
	return toCoreList(getDataAll), nil
}

func (repo *mysqlRoomRepository) GetData(id int) (data rooms.Core, err error) {
	var getData Rooms
	tx := repo.db.Preload("User").Preload("Categorys").First(&getData, id)
	if tx.Error != nil {
		return rooms.Core{}, tx.Error
	}
	return getData.toCore(), nil
}

func (repo *mysqlRoomRepository) UpdateData(id int, insert rooms.Core) (row int, err error) {
	tx := repo.db.Model(&Rooms{}).Where("id = ?", id).Updates(Rooms{ImageRoom: insert.ImageRoom, ImagePengelola: insert.ImagePengelola, Name: insert.Name, Capacity: insert.Capacity, RentalPrice: insert.RentalPrice, Address: insert.Address, City: insert.City, CategorysID: uint(insert.Categorys.ID)})
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRoomRepository) DeleteData(id int) (row int, err error) {
	var deleteData Rooms
	tx := repo.db.Unscoped().Delete(&deleteData, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *mysqlRoomRepository) GetToken(id int, idToken int) (data rooms.Core, err error) {
	var getData Rooms
	tx := repo.db.Preload("User").Preload("Categorys").First(&getData, id)
	if tx.Error != nil {
		return rooms.Core{}, tx.Error
	}
	if getData.toCore().User.ID != idToken {
		return rooms.Core{}, err
	}
	return getData.toCore(), nil
}

func (repo *mysqlRoomRepository) GetDataAllUserRoom(idToken int) (data []rooms.Core, err error) {
	var getDataUserRoom []Rooms
	tx := repo.db.Where("user_id = ?", idToken).Preload("User").Preload("Categorys").Find(&getDataUserRoom)
	if tx.Error != nil {
		return []rooms.Core{}, tx.Error
	}
	return toCoreList(getDataUserRoom), nil
}
