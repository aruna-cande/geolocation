package service

import (
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/aruna-cande/geolocation/pkg/geolocation/domain"
	"github.com/aruna-cande/geolocation/pkg/geolocation/service/mock"
)

func TestGeolocationDataService_GetGeolocationByIP(t *testing.T) {
	type test struct {
		ipAddress   string
		geolocation *domain.Geolocation
		error       error
	}

	geoData := domain.NewGeolocation("10.0.0.1", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346")
	tests := []test{
		{ipAddress: "10.0.0.1", geolocation: geoData, error: nil},
		{ipAddress: "10.1.0.0", geolocation: nil, error: sql.ErrNoRows},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepository(ctrl)

	for _, test := range tests {
		repo.EXPECT().GetGeolocationByIP(gomock.Any()).Return(test.geolocation, test.error)
		service := NewGeolocationDataService(repo)
		geoData, err := service.GetGeolocationByIP(test.ipAddress)

		assert.Equal(t, test.geolocation, geoData)
		assert.Equal(t, test.error, err)
	}
}
