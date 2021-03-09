package geolocation

type geolocation struct {
	id           string
	ipAddress    string
	countryCode  string
	country      string
	city         string
	latitude     float32
	longitude    float32
	mysteryValue int
}

func NewGeolocation(uuid string, countryCode string, country string, city string, latitude float32, longitude float32, mysteryValue int) {

}

func (g geolocation) GetUuid() string {
	return g.id
}

func (g geolocation) GetIpAddress() string {
	return g.ipAddress
}

func (g geolocation) GetCountryCode() string {
	return g.countryCode
}

func (g geolocation) GetCountry() string {
	return g.country
}

func (g geolocation) GetCity() string {
	return g.city
}

func (g geolocation) GetLatitude() float32 {
	return g.latitude
}

func (g geolocation) GetLongitude() float32 {
	return g.longitude
}

func (g geolocation) GetMysteryValue() int {
	return g.mysteryValue
}

