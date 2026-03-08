package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/aruna-cande/geolocation/cmd/geolocation-api/handler"
	"github.com/aruna-cande/geolocation/pkg/geolocation/adapters"
	"github.com/aruna-cande/geolocation/pkg/geolocation/service"
)

func main() {
	var logger = log.New(os.Stderr, "logger: ", log.Ldate|log.Ltime|log.Lshortfile)

	config, err := NewConfig()
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}

	dbInit := adapters.NewDBInitializer(config.PostgresUser, config.PostgresPassword, config.PostgresDb, config.PostgresHost, config.PostgresPort)
	db, err := dbInit.InitDatabase()
	if err != nil {
		logger.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	postgresRepository := adapters.NewGeolocationPostgresRepository(db)
	geolocationService := service.NewGeolocationDataService(postgresRepository)

	r := mux.NewRouter()

	n := negroni.New(
		negroni.HandlerFunc(cors),
		negroni.NewLogger(),
	)

	http.Handle("/", r)
	handler.CreateGeolocationHandler(r, *n, geolocationService)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf(":%d", config.APIPort),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err.Error())
	}
}

func cors(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "OPTIONS" {
		return
	}
	next(w, r)
}
