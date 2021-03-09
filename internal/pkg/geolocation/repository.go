package geolocation

import (
	"database/sql"
)

type PostgresRepository interface {
	AddGeolocation(geolocation geolocation) error
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

func (r *postgresRepository) AddGeolocation(geolocation geolocation) error {
	stmt, err := r.db.Prepare(`
		Insert INTO geolocation (id, ipAddress, countryCode, country, city, latitude, longitude, mysteryValue) 
		values($1,$2,$3,$4,$5,$6,$7,$8)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		geolocation.id,
		geolocation.ipAddress,
		geolocation.countryCode,
		geolocation.country,
		geolocation.city,
		geolocation.latitude,
		geolocation.longitude,
		geolocation.mysteryValue,
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
	stmt, err := r.db.Prepare(`SELECT id, ipAddress, countryCode, country, city, latitude, longitude, mysteryValue FROM geolocation WHERE ipAddress = $1`)
	if err != nil {
		return nil, err
	}
	var geoModel geolocation
	row := stmt.QueryRow(ipAddress)

	err = row.Scan( &geoModel.id, &geoModel.ipAddress, &geoModel.countryCode, &geoModel.country, &geoModel.city, &geoModel.latitude, geoModel.longitude, geoModel.mysteryValue)
	if err == sql.ErrNoRows{
		return nil, err
	}

	return &geoModel, nil
}
