package adapters

import (
	"Geolocation/pkg/geolocation/domain"
	"database/sql"
	"fmt"
	"strconv"
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
	for _, geolocation := range geolocations {
		values = append(values, fmt.Sprintf("(?, ?, ?, ?, ?, ?, ?)"))
		args = append(args, geolocation.IpAddress)
		args = append(args, geolocation.CountryCode)
		args = append(args, geolocation.Country)
		args = append(args, geolocation.City)
		args = append(args, geolocation.Latitude)
		args = append(args, geolocation.Longitude)
		args = append(args, geolocation.MysteryValue)
	}

	query := fmt.Sprintf(`INSERT INTO geolocations_data (IpAddress, CountryCode, Country, City, Latitude, Longitude, MysteryValue) Values %s`, strings.Join(values, ","))
	query = replacePattern(query, "?")
	stmt, err := r.db.Prepare(strings.TrimSuffix(query, ","))
	if err != nil {
		return int64(len(geolocations)), err
	}
	result, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	err = stmt.Close()
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		query = fmt.Sprintf("INSERT INTO geolocations_data (IpAddress, CountryCode, Country, City, Latitude, Longitude, MysteryValue) Values ($1, $2, $3, $4, $5, $6, $7)")
		for _, geolocation := range geolocations {
			stmt, err := r.db.Prepare(query)
			if err != nil {
				continue
			}

			_, err = stmt.Exec(geolocation.Id, geolocation.CountryCode, geolocation.Country, geolocation.City, geolocation.Latitude, geolocation.Longitude, geolocation.MysteryValue)
			if err != nil {
				continue
			}
			rowsAffected++
		}
		return 0, err
	}

	return rowsAffected, nil
}

func replacePattern(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

func (r *postgresRepository) GetGeolocationByIp(ipAddress string) (*domain.Geolocation, error) {
	var geolocation domain.Geolocation
	row := r.db.QueryRow("SELECT Id, IpAddress, CountryCode, Country, City, Latitude, Longitude, MysteryValue FROM geolocations_data WHERE IpAddress = $1", ipAddress)

	err := row.Scan(&geolocation.Id, &geolocation.IpAddress, &geolocation.CountryCode, &geolocation.Country, &geolocation.City, &geolocation.Latitude, geolocation.Longitude, geolocation.MysteryValue)
	if err == sql.ErrNoRows {
		return nil, err
	}

	return &geolocation, nil
}
