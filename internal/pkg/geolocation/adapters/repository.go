package adapters

import (
	"Geolocation/internal/pkg/geolocation/domain"
	"database/sql"
	"fmt"
	"strings"
)

type Repository interface {
	AddGeolocation(geolocations []*domain.Geolocation) (int64, error)
	GetGeolocationByIp(ipAddress string) (*domain.Geolocation, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewGeolocationPostgresRepository(
	pgClient *sql.DB,
) Repository {
	return &postgresRepository{
		db: pgClient,
	}
}

func (r *postgresRepository) AddGeolocation(geolocations []*domain.Geolocation) (int64, error) {
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
		Insert INTO geolocations_data (IpAddress, CountryCode, Country, City, Latitude, Longitude, MysteryValue) 
		values %s`, strings.Join(values, ",")))
	if err != nil {
		return int64(len(geolocations)), err
	}
	result, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}
	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *postgresRepository) GetGeolocationByIp(ipAddress string) (*domain.Geolocation, error) {
	stmt, err := r.db.Prepare(`SELECT Id, IpAddress, CountryCode, Country, City, Latitude, Longitude, MysteryValue FROM geolocations_data WHERE IpAddress = $1`)
	if err != nil {
		return nil, err
	}
	var geolocation domain.Geolocation
	row := stmt.QueryRow(ipAddress)

	err = row.Scan( &geolocation.Id, &geolocation.IpAddress, &geolocation.CountryCode, &geolocation.Country, &geolocation.City, &geolocation.Latitude, geolocation.Longitude, geolocation.MysteryValue)
	if err == sql.ErrNoRows{
		return nil, err
	}

	return &geolocation, nil
}
