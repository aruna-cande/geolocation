package domain

import (
	"context"
	"errors"
	"net"
	"strconv"
)

// Geolocation represents a geolocation record.
type Geolocation struct {
	ID           string
	IPAddress    string
	CountryCode  string
	Country      string
	City         string
	Latitude     float64
	Longitude    float64
	MysteryValue string
}

// NewGeolocation creates and validates a Geolocation from raw string inputs.
// Returns nil if any field is invalid.
func NewGeolocation(ipAddress, countryCode string, country string, city string, latitude string, longitude string, mysteryValue string) *Geolocation {
	floatLatitude, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return nil
	}

	floatLongitude, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return nil
	}

	geoData := &Geolocation{
		IPAddress:    ipAddress,
		CountryCode:  countryCode,
		Country:      country,
		City:         city,
		Latitude:     floatLatitude,
		Longitude:    floatLongitude,
		MysteryValue: mysteryValue,
	}
	err = geoData.Validate()
	if err != nil {
		return nil
	}

	return geoData
}

// Validate checks that all Geolocation fields contain sensible values.
func (g *Geolocation) Validate() error {
	if net.ParseIP(g.IPAddress) == nil || g.IPAddress == "" {
		return errors.New("invalid IP address")
	}
	if g.CountryCode == "" || g.Country == "" || g.City == "" {
		return errors.New("invalid location")
	}

	if g.Latitude >= 90 || g.Latitude <= -90 || g.Longitude >= 180 || g.Longitude <= -180 {
		return errors.New("invalid geographical coordinates")
	}

	if g.MysteryValue == "" {
		return errors.New("mystery value is missing")
	}
	return nil
}

// Repository defines the data-access contract for geolocation records.
// Implementations live in the adapters package.
type Repository interface {
	AddGeolocation(ctx context.Context, geolocations []*Geolocation) (int64, error)
	GetGeolocationByIP(ctx context.Context, ipAddress string) (*Geolocation, error)
}
