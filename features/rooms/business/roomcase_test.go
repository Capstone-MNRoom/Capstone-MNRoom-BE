package business

import (
	"be9/mnroom/features/rooms"
	"be9/mnroom/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestInsertData(t *testing.T) {
	repo := new(mocks.RoomData)
	insertData := rooms.Core{
		ID: 1,
		ImageRoom: "https://storage.googleapis.com/bucket/image-default.jpg",
		ImagePengelola: "https://storage.googleapis.com/bucket/image-default.jpg",
		RoomName: "Ruang Arjuna",
		HotelName: "Berlina",
		RentalPrice: 1000000,
		Capacity: 200,
		Address: "Jl. Raya Utama No. 123",
		City: "Surabaya",
		
	}


	t.Run("Success Insert", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(1, nil).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.InsertData(insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert to DB", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(0, errors.New("there is some error")).Once()
		srv := NewRoomBusiness(repo)
		res, err := srv.InsertData(insertData)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}