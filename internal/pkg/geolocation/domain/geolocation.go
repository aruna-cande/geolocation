package domain


import (
	"errors"
	"net"
	"strconv"
)

type Geolocation struct {
	Id           string
	IpAddress    string
	CountryCode  string
	Country      string
	City         string
	Latitude     float64
	Longitude    float64
	MysteryValue string
}

func NewGeolocation(ipAddress, countryCode string, country string, city string, latitude string, longitude string, mysteryValue string) *Geolocation {
	floatLatitude, err := strconv.ParseFloat(latitude,64)
	if err != nil {
		return nil
	}

	floatLongitude, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return nil
	}

	geoData := &Geolocation{
		IpAddress:    ipAddress,
		CountryCode:  countryCode,
		Country:      country,
		City:         city,
		Latitude:     floatLatitude,
		Longitude:    floatLongitude,
		MysteryValue: mysteryValue,
	}
	err = geoData.Validate()
	if err != nil{
		return nil
	}

	return geoData
}

func (g *Geolocation) Validate() error  {
	if net.ParseIP(g.IpAddress) == nil || g.IpAddress == "" {
		return errors.New("Invalid IP")
	}
	if  g.CountryCode  == "" || g.Country =="" || g.City==""  {
		return errors.New("invalid location")
	}

	if g.Latitude >=90 || g.Latitude <= -90 || g.Longitude >=180 || g.Longitude <= -180{
		return errors.New("Invalid geografical cordinates")
	}

	if g.MysteryValue == "" {
		return errors.New("Mystery value is missing")
	}
	return nil
}


