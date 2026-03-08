package service

import (
	"github.com/aruna-cande/geolocation/pkg/geolocation/domain"
)

// GeolocationDataService provides read access to geolocation records.
//
// Currently this service delegates directly to the repository. It exists as the
// natural extension point for cross-cutting concerns such as caching, input
// validation, rate-limiting, or metrics — without coupling those concerns to
// the handler or repository layers.
type GeolocationDataService interface {
	GetGeolocationByIP(ipAddress string) (*domain.Geolocation, error)
}

type geolocationDataService struct {
	repo domain.Repository
}

// NewGeolocationDataService creates a new GeolocationDataService.
func NewGeolocationDataService(repository domain.Repository) GeolocationDataService {
	return &geolocationDataService{repository}
}

func (s *geolocationDataService) GetGeolocationByIP(ipAddress string) (*domain.Geolocation, error) {
	gsData, err := s.repo.GetGeolocationByIP(ipAddress)
	if err != nil {
		return nil, err
	}

	return gsData, nil
}
