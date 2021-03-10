package geolocation

import (
	"database/sql"
)

type PostgresRepository interface {
	AddGeolocation(geolocation *geolocation) error
	GetGeolocationByIp(ipAddress string) (*geolocation, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewGeolocationFirestoreRepository(
	pgClient *sql.DB,
) PostgresRepository {
	return &postgresRepository{
		db: pgClient,
	}
}

func (r *postgresRepository) AddGeolocation(geolocation *geolocation) error {
	stmt, err := r.db.Prepare(`
		Insert INTO geolocation (Id, IpAddress, CountryCode, Country, City, Latitude, Longitude, MysteryValue) 
		values($1,$2,$3,$4,$5,$6,$7,$8)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		geolocation.Id,
		geolocation.IpAddress,
		geolocation.CountryCode,
		geolocation.Country,
		geolocation.City,
		geolocation.Latitude,
		geolocation.Longitude,
		geolocation.MysteryValue,
	)
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
