package service

import (
	"github.com/aruna-cande/geolocation/pkg/geolocation/adapters"
	"github.com/aruna-cande/geolocation/pkg/geolocation/domain"
)

// GeolocationDataService provides read access to geolocation records.
type GeolocationDataService interface {
	GetGeolocationByIP(ipAddress string) (*domain.Geolocation, error)
}

type geolocationDataService struct {
	fr adapters.Repository
}

// NewGeolocationDataService creates a new GeolocationDataService.
func NewGeolocationDataService(repository adapters.Repository) GeolocationDataService {
	return &geolocationDataService{repository}
}

func (s *geolocationDataService) GetGeolocationByIP(ipAddress string) (*domain.Geolocation, error) {
	gsData, err := s.fr.GetGeolocationByIP(ipAddress)
	if err != nil {
		return nil, err
	}

	return gsData, nil
}
