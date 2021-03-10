package geolocation

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

type Service interface {
	ImportGeolocationData()
	GetGeolocationByIp(ipAdrress string) (*geolocation, error)
}

type service struct {
	fr Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) ImportGeolocationData() {

	r := csv.NewReader(strings.NewReader("data_dump.csv"))

	var locations []*geolocation
	var invalidRecords int
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		ipAddress := record[0]
		countryCode := record[2]
		country := record[2]
		city := record[3]
		latitude := record[4]
		longitude := record[5]
		mysteryValue := record[6]
		geoData := NewGeolocation(ipAddress, countryCode, country, city, latitude, longitude, mysteryValue)

		if geoData == nil {
			invalidRecords++
		}
		locations = append(locations, geoData)
		fmt.Println(record)
	}
	locationChunks := getChunksOfGeolocationData(locations, 1000)
	// i := 0; i < len(locationChunks); i++
	for _, chunk := range locationChunks{
		err := s.fr.AddGeolocation(chunk)

		if err == nil {
			fmt.Println("failed to add locations")
		}
	}

	//if err != nil
	fmt.Println("Failed ", err.Error())
}

func getChunksOfGeolocationData(locations []*geolocation, chunkSize int) [][]*geolocation {
	var locationChunks [][]*geolocation

	for i := 0; i < len(locations); i += chunkSize {
		end := i + chunkSize

		if end > len(locations) {
			end = len(locations)
		}
		locationChunks = append(locationChunks, locations[i:end])
	}

	return locationChunks
}

func (s *service) GetGeolocationByIp(ipAddress string) (*geolocation, error) {
	gsData, err := s.fr.GetGeolocationByIp(ipAddress)

	if err != nil {
		return nil, err
	}

	return gsData, nil

}
