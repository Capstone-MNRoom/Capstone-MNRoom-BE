package business

import (
	"be9/mnroom/features/facilitys"
	"be9/mnroom/features/roomfacilitys"
	"be9/mnroom/features/rooms"
	"be9/mnroom/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetData(t *testing.T) {
	repo := new(mocks.RoomFacilityData)
	list := []roomfacilitys.Core{
		{
			ID:        1,
			Rooms:     rooms.Core{ID: 1, RoomName: "Ruang Arjuna"},
			Facilitys: facilitys.Core{ID: 1, Name: "Projector"},
		},
	}

	t.Run("Success Get Data", func(t *testing.T) {
		repo.On("GetData", mock.Anything, 1).Return(list, nil).Once()
		srv := NewRoomFacilityBusiness(repo)

		res, err := srv.GetData(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, "Projector", res[0].Facilitys.Name)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data", func(t *testing.T) {
		repo.On("GetData", mock.Anything, 1).Return(nil, errors.New("not found")).Once()
		srv := NewRoomFacilityBusiness(repo)

		res, err := srv.GetData(context.Background(), 1)
		assert.Error(t, err)
		assert.Nil(t, res)
		repo.AssertExpectations(t)
	})
}

func TestGetDataRow(t *testing.T) {
	repo := new(mocks.RoomFacilityData)

	t.Run("Success Get Data Row", func(t *testing.T) {
		repo.On("GetDataRow", mock.Anything, 1).Return(1, nil).Once()
		srv := NewRoomFacilityBusiness(repo)

		res, err := srv.GetDataRow(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error Get Data Row", func(t *testing.T) {
		repo.On("GetDataRow", mock.Anything, 1).Return(0, errors.New("not found")).Once()
		srv := NewRoomFacilityBusiness(repo)

		res, err := srv.GetDataRow(context.Background(), 1)
		assert.Error(t, err)
		assert.Equal(t, 0, res)
		repo.AssertExpectations(t)
	})
}
