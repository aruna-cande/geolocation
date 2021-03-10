package geolocation

import (
	"errors"
	"strconv"
)

type geolocation struct {
	Id           string
	IpAddress    string
	CountryCode  string
	Country      string
	City         string
	Latitude     float64
	Longitude    float64
	MysteryValue int
}

func NewGeolocation(ipAddress, countryCode string, country string, city string, latitude string, longitude string, mysteryValue int) *geolocation {
	flatitude, err := strconv.ParseFloat(latitude,64)
	if err != nil {
		return nil
	}

	flongitude, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return nil
	}

	geoData := &geolocation{
		IpAddress:    ipAddress,
		CountryCode:  countryCode,
		Country:      country,
		City:         city,
		Latitude:     flatitude,
		Longitude:    flongitude,
		MysteryValue: mysteryValue,
	}
	err = geoData.Validate()
	if err != nil{
		return nil
	}

	return geoData
}

func (g * geolocation) Validate() error  {

	if g.IpAddress == "" || g.CountryCode  == "" || g.Country =="" || g.City=="" || g.Latitude >=90 || g.Latitude <= -90 || g.Longitude >=180 || g.Longitude <= -180 {
		return errors.New("invalid location")
	}
	return nil
}

