package geolocation

import (
	"fmt"
)

type Service interface {
	ImportGeolocationData(geolocation geolocation)
	GetGeolocationByIp(ipAdrress string)(*geolocation, error)
}

type service struct {
	fr PostgresRepository
}

func NewService(repository PostgresRepository) Service  {
	return &service{repository}
}

func (s *service) ImportGeolocationData(geolocation geolocation) {
	err := s.fr.AddGeolocation(geolocation)

	//if err != nil
	fmt.Println("Failed ", err.Error())
}

func (s *service) GetGeolocationByIp(ipAddress string) (*geolocation, error){
	gsData, err := s.fr.GetGeolocationByIp(ipAddress)

	if err != nil{
		return nil, err
	}

	return gsData, nil

}

