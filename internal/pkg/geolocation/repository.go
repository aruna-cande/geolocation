package geolocation

import (
	"database/sql"
	"fmt"
	"strings"
)

type Repository interface {
	AddGeolocation(geolocations []*geolocation) error
	GetGeolocationByIp(ipAddress string) (*geolocation, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewGeolocationFirestoreRepository(
	pgClient *sql.DB,
) Repository {
	return &postgresRepository{
		db: pgClient,
	}
}

func (r *postgresRepository) AddGeolocation(geolocations []*geolocation) error {
	var values []string
	var args []interface{}
	for  _, geolocation := range geolocations {
		values = append(values, "(?, ?, ?, ?, ?, ?, ?)")
		args = append(args, geolocation.IpAddress)
		args = append(args, geolocation.CountryCode)
		args = append(args, geolocation.Country)
		args = append(args, geolocation.City)
		args = append(args, geolocation.Latitude)
		args = append(args, geolocation.Longitude)
		args = append(args, geolocation.MysteryValue)
	}

	stmt, err := r.db.Prepare(fmt.Sprintf(`
		Insert INTO geolocation (IpAddress, CountryCode, Country, City, Latitude, Longitude, MysteryValue) 
		values %s`, strings.Join(values, ",")))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(args)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresRepository) GetGeolocationByIp(ipAddress string) (*geolocation, error) {
	stmt, err := r.db.Prepare(`SELECT Id, IpAddress, CountryCode, Country, City, Latitude, Longitude, MysteryValue FROM geolocation WHERE IpAddress = $1`)
	if err != nil {
		return nil, err
	}
	var geoModel geolocation
	row := stmt.QueryRow(ipAddress)

	err = row.Scan( &geoModel.Id, &geoModel.IpAddress, &geoModel.CountryCode, &geoModel.Country, &geoModel.City, &geoModel.Latitude, geoModel.Longitude, geoModel.MysteryValue)
	if err == sql.ErrNoRows{
		return nil, err
	}

	return &geoModel, nil
}
