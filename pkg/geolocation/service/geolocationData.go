package service

import (
	"geolocation/pkg/geolocation/adapters"
	"geolocation/pkg/geolocation/domain"
)

type GeolocationDataService interface {
	GetGeolocationByIp(ipAddress string) (*domain.Geolocation, error)
}

type geolocationDataService struct {
	fr adapters.Repository
}

func NewGeolocationDataService(repository adapters.Repository) GeolocationDataService {
	return &geolocationDataService{repository}
}

func (s *geolocationDataService) GetGeolocationByIp(ipAddress string) (*domain.Geolocation, error) {
	gsData, err := s.fr.GetGeolocationByIp(ipAddress)

	if err != nil {
		return nil, err
	}

	return gsData, nil

}
