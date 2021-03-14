package adapters

import (
	"Geolocation/pkg/geolocation/domain"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func newFixtureGeolocations() []*domain.Geolocation {
	var geolocations []*domain.Geolocation
	geoData := &domain.Geolocation{
		Id:           "123e4567-e89b-12d3-a456-426614174000",
		IpAddress:    "200.106.141.15",
		CountryCode:  "SI",
		Country:      "Nepal",
		City:         "DuBuquemouth",
		Latitude:     -84.87503094689836,
		Longitude:    7.206435933364332,
		MysteryValue: "7823011346",
	}
	geoData2 := &domain.Geolocation{
		Id:           "123e4567-e89b-12d3-a456-426614175000",
		IpAddress:    "160.103.7.140",
		CountryCode:  "NI",
		Country:      "Nicaragua",
		City:         "New Neva",
		Latitude:     -68.31023296602508,
		Longitude:    -37.62435199624531,
		MysteryValue: "7301823115",
	}
	geolocations = append(geolocations, geoData, geoData2)
	return geolocations
}

func TestPostgresRepository_AddGeolocation(t *testing.T) {
	geolocations := newFixtureGeolocations()

	db, mock, err := sqlmock.New()
	if err != nil {
		logrus.Error(err.Error())
	}
	repo := NewGeolocationPostgresRepository(db)

	query := "INSERT INTO geolocations_data (IpAddress, CountryCode, Country, City, Latitude, Longitude, MysteryValue) Values ($1, $2, $3, $4, $5, $6, $7),($8, $9, $10, $11, $12, $13, $14)"
	query = regexp.QuoteMeta(query)
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(geolocations[0].IpAddress, geolocations[0].CountryCode, geolocations[0].Country,
		geolocations[0].City, geolocations[0].Latitude, geolocations[0].Longitude, geolocations[0].MysteryValue,
		geolocations[1].IpAddress, geolocations[1].CountryCode, geolocations[1].Country,
		geolocations[1].City, geolocations[1].Latitude, geolocations[1].Longitude, geolocations[1].MysteryValue).WillReturnResult(sqlmock.NewResult(1, 2))

	_, err = repo.AddGeolocation(geolocations)
	assert.NoError(t, err)
}

func TestPostgresRepository_GetGeolocationByIp(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		logrus.Error(err.Error())
	}
	repo := NewGeolocationPostgresRepository(db)
	geolocations := newFixtureGeolocations()

	query := "SELECT Id, IpAddress, CountryCode, Country, City, Latitude, Longitude, MysteryValue FROM geolocations_data WHERE IpAddress = $1"
	row := sqlmock.NewRows([]string{"Id", "IpAddress", "CountryCode", "Country", "City", "Latitude", "Longitude", "MysteryValue"}).
		AddRow(geolocations[0].Id, geolocations[0].IpAddress, geolocations[0].CountryCode, geolocations[0].Country,
			geolocations[0].City, geolocations[0].Latitude, geolocations[0].Longitude, geolocations[0].MysteryValue)
	mock.ExpectQuery(query).WithArgs(geolocations[0].IpAddress).WillReturnRows(row)

	geolocation, err := repo.GetGeolocationByIp(geolocations[0].IpAddress)
	assert.NotNil(t, geolocation)
	assert.Nil(t, err)
}
