package service

import (
	"Geolocation/pkg/geolocation/adapters"
	"Geolocation/pkg/geolocation/domain"
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"
)

type ImporterService interface {
	ImportGeolocationData(filepath string) (Statistics, error)
}

type importerService struct {
	fr adapters.Repository
	log *log.Logger
}

func NewImporterService(repository adapters.Repository, logger *log.Logger) ImporterService {
	return &importerService{repository, logger}
}

func (s *importerService) ImportGeolocationData(filepath string) (Statistics, error) {
	started := time.Now()
	s.log.Println("inside ImportGeolocationData")
	data, err := os.Open(filepath)
	if err != nil {
		s.log.Println("File "+filepath+" not found")
		return Statistics{}, err
	}
	r := csv.NewReader(data)

	keys := make(map[string]bool)
	var locations []*domain.Geolocation
	var discarded int64
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			s.log.Fatal(err)
			return Statistics{}, err
		}
		if isImporterCsvHeader(record) {
			continue
		}

		ipAddress := record[0]
		countryCode := record[1]
		country := record[2]
		city := record[3]
		latitude := record[4]
		longitude := record[5]
		mysteryValue := record[6]

		if _, value := keys[ipAddress]; !value {
			keys[ipAddress] = true
			geoData := domain.NewGeolocation(ipAddress, countryCode, country, city, latitude, longitude, mysteryValue)

			if geoData == nil {
				discarded++
				continue
			}
			locations = append(locations, geoData)
			s.log.Print(record)
		} else {
			discarded++
		}
	}
	locationChunks := getChunksOfGeolocationData(locations, 500)

	var discardedDb int64
	for _, chunk := range locationChunks {
		rowsAffected, err := s.fr.AddGeolocation(chunk)
		if err != nil {
			s.log.Println("Failed with error: " + err.Error())
			discardedDb = discardedDb + int64(len(chunk))
		}else{
			discardedDb = discardedDb + (int64(len(chunk)) - rowsAffected)
		}
	}
	duration := time.Since(started)
	accepted := int64(len(locations)) - discardedDb
	statistics := Statistics{
		TimeElapsed: duration,
		Accepted:    accepted,
		Discarded:   discarded + discardedDb,
	}
	return statistics, nil
}

func isImporterCsvHeader(record []string) bool{
	if record[0] == "ip_address" && record[1] == "country_code" &&
		record[2] == "country" && record[3] == "city" &&
		record[4] == "latitude" && record[5] == "longitude" &&
		record[6] == "mystery_value"{
		return true
	}
	return false
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
