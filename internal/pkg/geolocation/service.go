package geolocation

import (
	"context"
	"fmt"
)

type Service interface {
	ImportGeolocationData(ctx context.Context, geolocation geolocation)
	GetGeolocationByIp(ctx context.Context, ipAdrress string)(*geolocation, error)
}

type service struct {
	fr FirestoreRepository
}

func newService(repository FirestoreRepository) Service  {
	return &service{repository}
}

func (s *service) ImportGeolocationData(ctx context.Context, geolocation geolocation) {
	err := s.fr.AddGeolocation(ctx, geolocation)

	//if err != nil
	fmt.Println("Failed ", err.Error())
}

func (s *service) GetGeolocationByIp(ctx context.Context, ipAdrress string) (*geolocation, error){
	gsData, err := s.fr.GetGeolocationByIp(ctx, ipAdrress)

	if err != nil{
		return nil, err
	}

	return gsData, nil

}

