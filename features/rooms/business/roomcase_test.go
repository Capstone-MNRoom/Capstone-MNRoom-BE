package business

import (
	"be9/mnroom/features/categorys"
	"be9/mnroom/features/facilitys"
	"be9/mnroom/features/rooms"
	"be9/mnroom/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertData(t *testing.T) {
	repo := new(mocks.RoomData)
	insertData := rooms.Core{
		ID:             1,
		ImageRoom:      "https://storage.googleapis.com/bucket/image-default.jpg",
		ImagePengelola: "https://storage.googleapis.com/bucket/image-default.jpg",
		RoomName:       "Ruang Arjuna",
		HotelName:      "Berlina",
		RentalPrice:    1000000,
		Capacity:       200,
		Address:        "Jl. Raya Utama No. 123",
		City:           "Surabaya",
	}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("InsertData", mock.Anything, mock.Anything).Return(insertData, nil).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.InsertData(context.Background(), insertData)
		assert.NoError(t, err)
		assert.Equal(t, insertData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		repo.On("InsertData", mock.Anything, mock.Anything).Return(rooms.Core{}, errors.New("there is some error")).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.InsertData(context.Background(), insertData)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, rooms.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestInsertDataRoomFacilitys(t *testing.T) {
	repo := new(mocks.RoomData)
	insertData := rooms.CoreRoomFacilitys{
		Rooms: rooms.Core{
			ID: 1,
		},
		Facilitys: facilitys.Core{
			ID: 1,
		},
	}

	t.Run("Success Insert Room Facility", func(t *testing.T) {
		repo.On("InsertDataRoomFacilitys", mock.Anything, insertData).Return(1, nil).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.InsertDataRoomFacilitys(context.Background(), insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Insert Room Facility", func(t *testing.T) {
		repo.On("InsertDataRoomFacilitys", mock.Anything, insertData).Return(0, errors.New("insert failed")).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.InsertDataRoomFacilitys(context.Background(), insertData)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}

func TestGetDataAll(t *testing.T) {
	repo := new(mocks.RoomData)
	roomList := []rooms.Core{
		{
			ID:          1,
			RoomName:    "Ruang Arjuna",
			RentalPrice: 1000000,
		},
	}

	t.Run("Success Get Data All", func(t *testing.T) {
		repo.On("GetDataAll", mock.Anything, 0).Return(roomList, nil).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetDataAll(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, "Ruang Arjuna", res[0].RoomName)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data All", func(t *testing.T) {
		repo.On("GetDataAll", mock.Anything, 0).Return(nil, errors.New("query error")).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetDataAll(context.Background(), 1)
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetData(t *testing.T) {
	repo := new(mocks.RoomData)
	roomData := rooms.Core{
		ID:          1,
		RoomName:    "Ruang Arjuna",
		RentalPrice: 1000000,
	}

	t.Run("Success Get Data", func(t *testing.T) {
		repo.On("GetData", mock.Anything, 1).Return(roomData, nil).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetData(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, roomData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data", func(t *testing.T) {
		repo.On("GetData", mock.Anything, 1).Return(rooms.Core{}, errors.New("not found")).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetData(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, rooms.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetDataIDRoom(t *testing.T) {
	repo := new(mocks.RoomData)

	t.Run("Success Get Data ID Room", func(t *testing.T) {
		repo.On("GetDataIDRoom", mock.Anything, 1).Return(1, nil).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetDataIDRoom(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data ID Room", func(t *testing.T) {
		repo.On("GetDataIDRoom", mock.Anything, 1).Return(0, errors.New("not found")).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetDataIDRoom(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateData(t *testing.T) {
	repo := new(mocks.RoomData)
	updateData := rooms.Core{
		RoomName:    "Ruang Bima",
		RentalPrice: 1500000,
	}

	t.Run("Success Update", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, 1, updateData).Return(1, nil).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.UpdateData(context.Background(), 1, updateData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Update", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, 1, updateData).Return(0, errors.New("update error")).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.UpdateData(context.Background(), 1, updateData)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T) {
	repo := new(mocks.RoomData)

	t.Run("Success Delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, 1).Return(1, nil).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.DeleteData(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, 1).Return(0, errors.New("delete error")).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.DeleteData(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}

func TestGetToken(t *testing.T) {
	repo := new(mocks.RoomData)
	roomData := rooms.Core{
		ID:       1,
		RoomName: "Ruang Arjuna",
	}

	t.Run("Success Get Token", func(t *testing.T) {
		repo.On("GetToken", mock.Anything, 1, 10).Return(roomData, nil).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetToken(context.Background(), 1, 10)
		assert.NoError(t, err)
		assert.Equal(t, roomData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Token", func(t *testing.T) {
		repo.On("GetToken", mock.Anything, 1, 10).Return(rooms.Core{}, errors.New("unauthorized")).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetToken(context.Background(), 1, 10)
		assert.Error(t, err)
		assert.Equal(t, rooms.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestGetDataAllUserRoom(t *testing.T) {
	repo := new(mocks.RoomData)
	roomList := []rooms.Core{
		{ID: 1, RoomName: "Ruang Arjuna"},
		{ID: 2, RoomName: "Ruang Nakula"},
	}

	t.Run("Success Get Data All User Room", func(t *testing.T) {
		repo.On("GetDataAllUserRoom", mock.Anything, 10).Return(roomList, nil).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetDataAllUserRoom(context.Background(), 10)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(res))
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data All User Room", func(t *testing.T) {
		repo.On("GetDataAllUserRoom", mock.Anything, 10).Return(nil, errors.New("query error")).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetDataAllUserRoom(context.Background(), 10)
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetDataByCategory(t *testing.T) {
	repo := new(mocks.RoomData)
	roomList := []rooms.Core{
		{
			ID:       1,
			RoomName: "Meeting Room A",
			Categorys: categorys.Core{ID: 2, CategoryName: "Meeting"},
		},
	}

	t.Run("Success Get Data By Category", func(t *testing.T) {
		repo.On("GetDataByCategory", mock.Anything, 2).Return(roomList, nil).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetDataByCategory(context.Background(), 2)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, "Meeting Room A", res[0].RoomName)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data By Category", func(t *testing.T) {
		repo.On("GetDataByCategory", mock.Anything, 2).Return(nil, errors.New("not found")).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.GetDataByCategory(context.Background(), 2)
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}
