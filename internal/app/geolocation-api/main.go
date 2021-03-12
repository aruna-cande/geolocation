package main

import (
	"Geolocation/internal/app/geolocation-api/handler"
	"Geolocation/internal/pkg/geolocation/service"
	"Geolocation/internal/pkg/geolocation/adapters"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	//initDatabase(user, password, dbname)
	initDb := adapters.NewInitDb(user, password, dbname)
	db := initDb.InitDatabase()

	/*connectionString := getConnectionString(user, password, dbname)
	db, err := sql.Open("postgres", connectionString)
	defer db.Close()

	createTableGeolocations(db)

	if err != nil {
		fmt.Sprintln("erro", err)
	}*/

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

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
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
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
