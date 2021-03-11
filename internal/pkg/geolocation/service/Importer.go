package service

import (
	"Geolocation/internal/pkg/geolocation/adapters"
	"Geolocation/internal/pkg/geolocation/domain"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ImporterService interface {
	ImportGeolocationData(filepath string) (Statistics, error)
}

type importerService struct {
	fr adapters.Repository
}

func NewImporterService(repository adapters.Repository) ImporterService {
	return &importerService{repository}
}

var (
	Client adapters.HTTPClient
)

func init() {
	Client = &http.Client{}
}

func (s *importerService) ImportGeolocationData(filepath string) (Statistics, error) {
	started := time.Now()

	data, err := os.Open(filepath)
	if err != nil {
		return Statistics{}, err
	}
	r := csv.NewReader(data)

	var locations []*domain.Geolocation
	var discarded int64
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return Statistics{}, err
		}
		if record[0] == "ip_address" {
			continue
		}
		ipAddress := record[0]
		countryCode := record[1]
		country := record[2]
		city := record[3]
		latitude := record[4]
		longitude := record[5]
		mysteryValue := record[6]
		geoData := domain.NewGeolocation(ipAddress, countryCode, country, city, latitude, longitude, mysteryValue)

		if geoData == nil {
			discarded++
			continue
		}
		locations = append(locations, geoData)
		fmt.Println(record)
	}
	locationChunks := getChunksOfGeolocationData(locations, 1000)
	// i := 0; i < len(locationChunks); i++
	for _, chunk := range locationChunks {
		rowsAffected, err := s.fr.AddGeolocation(chunk)
		if err != nil {
			fmt.Println("failed to add locations")
			discarded = discarded + int64(len(chunk))
		}
		discarded = discarded + (int64(len(chunk)) - rowsAffected)
	}
	duration := time.Since(started)
	accepted := int64(len(locations)) - discarded
	statistics := Statistics{
		TimeElapsed: duration,
		Accepted: accepted,
		Discarded:   discarded,
	}
	return statistics, nil
}

func DownloadCsvFile(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return Client.Do(request)
}

func getChunksOfGeolocationData(locations []*domain.Geolocation, chunkSize int) [][]*domain.Geolocation {
	var locationChunks [][]*domain.Geolocation

	for i := 0; i < len(locations); i += chunkSize {
		end := i + chunkSize

		if end > len(locations) {
			end = len(locations)
		}
		locationChunks = append(locationChunks, locations[i:end])
	}

	return locationChunks
}
