package handler

import (
	"Geolocation/internal/pkg/geolocation"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)


func CreateGeolocationHandler(r *mux.Router, n negroni.Negroni, service geolocation.Service){
	r.Handle("/api/geolocations", n.With(
		negroni.Wrap(GetGeolocationByIp(service)),
	)).Methods("GET", "OPTIONS").Name("GetGeolocationByIp")
}

func GetGeolocationByIp(service geolocation.Service) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAddress := r.URL.Query().Get("ipaddress")
		data, err := service.GetGeolocationByIp(ipAddress)

		fmt.Sprintln("data value: ", data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		geoData := &geolocationResult{
			ipAddress:   data.GetIpAddress(),
			countryCode: data.GetCountryCode(),
			country:     data.GetCountry(),
			city:        data.GetCity(),
			latitude:    data.GetLatitude(),
			longitude:   data.GetLongitude(),
		}

		if err := json.NewEncoder(w).Encode(geoData); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}