package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
	"go.uber.org/mock/gomock"

	"github.com/aruna-cande/geolocation/pkg/geolocation/domain"
	"github.com/aruna-cande/geolocation/pkg/geolocation/service/mock"
)

func TestGetGeolocationByIP(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockGeolocationDataService(controller)
	r := mux.NewRouter()
	n := negroni.New()
	CreateGeolocationHandler(r, *n, service)
	path, err := r.GetRoute("GetGeolocationByIP").GetPathTemplate()
	assert.Nil(t, err)
	assert.Equal(t, "/api/geolocations", path)
	geoData := domain.NewGeolocation("10.0.0.1", "SI", "Nepal", "DuBuquemouth", "-84.87503094689836", "7.206435933364332", "7823011346")
	service.EXPECT().
		GetGeolocationByIP(gomock.Any(), geoData.IPAddress).
		Return(geoData, nil)
	handler := GetGeolocationByIP(service)
	r.Handle("/api/geolocations", handler)
	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/api/geolocations?ipaddress=" + geoData.IPAddress)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	var data *geolocationResult
	json.NewDecoder(res.Body).Decode(&data)
	assert.NotNil(t, data)
	assert.Equal(t, geoData.IPAddress, data.IPAddress)
}
