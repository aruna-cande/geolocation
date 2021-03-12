package handler

type geolocationResult struct {
	IpAddress string `json:"ipAddress"`
	CountryCode string `json:"countryCode"`
	Country string `json:"country"`
	City string `json:"city"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}