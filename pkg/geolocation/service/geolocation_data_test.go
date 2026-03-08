package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/aruna-cande/geolocation/pkg/geolocation/domain"
	"github.com/aruna-cande/geolocation/pkg/geolocation/service/mock"
)

func TestGeolocationDataService_GetGeolocationByIP(t *testing.T) {
	geoData := domain.NewGeolocation("10.0.0.1", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346")

	tests := []struct {
		name        string
		ipAddress   string
		geolocation *domain.Geolocation
		err         error
	}{
		{name: "Found", ipAddress: "10.0.0.1", geolocation: geoData, err: nil},
		{name: "NotFound", ipAddress: "10.1.0.0", geolocation: nil, err: sql.ErrNoRows},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepository(ctrl)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo.EXPECT().GetGeolocationByIP(gomock.Any(), gomock.Any()).Return(tc.geolocation, tc.err)
			service := NewGeolocationDataService(repo)
			geoData, err := service.GetGeolocationByIP(context.Background(), tc.ipAddress)

			assert.Equal(t, tc.geolocation, geoData)
			assert.Equal(t, tc.err, err)
		})
	}
}
