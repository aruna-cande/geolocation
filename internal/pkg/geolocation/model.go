package geolocation

type geolocation struct {
	uuid string
	ipAddress string
	countryCode string
	country string
	city string
	latitude float32
	longitude float32
	mysteryValue int
}

func NewGeolocation(uuid string, countryCode string, country string, city string, latitude float32, longitude float32, mysteryValue int) {

}

