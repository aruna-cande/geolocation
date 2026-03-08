package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/aruna-cande/geolocation/pkg/geolocation/service"
)

// CreateGeolocationHandler registers the geolocation routes on the router.
func CreateGeolocationHandler(r *mux.Router, n negroni.Negroni, service service.GeolocationDataService) {
	r.Handle("/api/geolocations", n.With(
		negroni.Wrap(GetGeolocationByIP(service)),
	)).Methods("GET", "OPTIONS").Name("GetGeolocationByIP")
}

// GetGeolocationByIP returns an HTTP handler that looks up geolocation by IP.
func GetGeolocationByIP(service service.GeolocationDataService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ipAddress := r.URL.Query().Get("ipaddress")
		data, err := service.GetGeolocationByIP(r.Context(), ipAddress)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if data == nil {
			http.Error(w, "geolocation not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		geoData := &geolocationResult{
			IPAddress:   data.IPAddress,
			CountryCode: data.CountryCode,
			Country:     data.Country,
			City:        data.City,
			Latitude:    data.Latitude,
			Longitude:   data.Longitude,
		}
		if err := json.NewEncoder(w).Encode(geoData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
