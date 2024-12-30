package application_test

import (
	"testing"

	"github.com/codeedu/go-hexagonal/application"
	mock_application "github.com/codeedu/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistance := mock_application.NewMockProductPersistanceInterface(ctrl)
	persistance.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistance: persistance,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}
