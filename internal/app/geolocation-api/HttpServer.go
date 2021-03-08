package geolocation_api

import (
	"context"
	"Geolocation/internal/pkg/geolocation"
	"fmt"
	"net/http"
)

type HttpServer struct {
	gs geolocation.Service
}

func NewHttpServer(gs geolocation.Service) HttpServer{
	return HttpServer{gs}
}

func (hs HttpServer) GetGeolocationByIp(w http.ResponseWriter, r *http.Request) http.ConnState{
	ctx := r.Context()
	data, err := hs.gs.GetGeolocationByIp(ctx, "1234")

	fmt.Sprintln("data value: ", data)
	if err != nil{
		return http.StatusNotFound
	}

	return http.StatusOK
}