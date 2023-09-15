package rocket

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("test get rocket by id", func(t *testing.T) {
		rockerStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rockerStoreMock.EXPECT().GetRocketByID(id).Return(Rocket{
			ID: id,
		}, nil)

		rocketService := New(rockerStoreMock)
		rkt, err := rocketService.GetRocketByID(
			context.Background(),
			id,
		)
		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rkt.ID)
	})

	t.Run("test insert rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.EXPECT().InsertRocket(Rocket{
			ID: id,
		}).Return(Rocket{
			ID: id,
		}, nil)

		rockerService := New(rocketStoreMock)
		rkt, err := rockerService.InsertRocket(
			context.Background(),
			Rocket{
				ID: id,
			},
		)

		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rkt.ID)
	})

	t.Run("test delete rocket", func(t *testing.T) {
		rockerStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rockerStoreMock.EXPECT().DeleteRocket(id).Return(nil)
		rocketService := New(rockerStoreMock)
		err := rocketService.DeleteRocket(id)
		assert.NoError(t, err)
	})
}
