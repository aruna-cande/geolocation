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
	MysteryValue int64
}

func NewGeolocation(ipAddress, countryCode string, country string, city string, latitude string, longitude string, mysteryValue string) *geolocation {
	floatLatitude, err := strconv.ParseFloat(latitude,64)
	if err != nil {
		return nil
	}

	floatLongitude, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return nil
	}

	intMysteryValue, err := strconv.ParseInt(mysteryValue,10, 64)

	geoData := &geolocation{
		IpAddress:    ipAddress,
		CountryCode:  countryCode,
		Country:      country,
		City:         city,
		Latitude:     floatLatitude,
		Longitude:    floatLongitude,
		MysteryValue: intMysteryValue,
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

