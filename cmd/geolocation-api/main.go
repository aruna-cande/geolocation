package main

import (
	"Geolocation/cmd/geolocation-api/handler"
	"Geolocation/pkg/geolocation/adapters"
	"Geolocation/pkg/geolocation/service"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
	"os"
	"strconv"
	"time"
	"log"
)

func main() {
	var logger = log.New(os.Stderr, "logger: ", log.Ldate|log.Ltime|log.Lshortfile)

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port, err := strconv.ParseInt(os.Getenv("POSTGRES_PORT"),10,64)
	if err != nil{
		logger.Fatal(err.Error())
		return
	}

	initDb := adapters.NewInitDb(user, password, dbname, host, port)
	db := initDb.InitDatabase()

	postgresRepository := adapters.NewGeolocationPostgresRepository(db)
	geolocationService := service.NewGeolocationDataService(postgresRepository)

	r := mux.NewRouter()

	n := negroni.New(
		negroni.HandlerFunc(Cors),
		negroni.NewLogger(),
	)

	http.Handle("/", r)
	handler.CreateGeolocationHandler(r, *n, geolocationService)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + os.Getenv("API_PORT"),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatal(err.Error())
	}
}

func Cors(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		return
	}
	next(w, r)
}
