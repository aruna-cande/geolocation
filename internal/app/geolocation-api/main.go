package main

import (
	"Geolocation/internal/app/geolocation-api/handler"
	"Geolocation/internal/pkg/geolocation"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
)

func main() {
	connectionString := getConnectionString()
	db, err := sql.Open("postgres", connectionString)
	createTableGeolocation(db)

	if err != nil {
		fmt.Sprintln("erro", err)
	}

	postgresRepository := geolocation.NewGeolocationFirestoreRepository(db)
	geolocationService := geolocation.NewService(postgresRepository)

	r := mux.NewRouter()

	n := negroni.New(
		negroni.HandlerFunc(Cors),
		negroni.NewLogger(),
	)

	http.Handle("/", r)
	handler.CreateGeolocationHandler(r, *n, geolocationService)
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	srv := &http.Server{
		//ReadTimeout:  5 * time.Second,
		//WriteTimeout: 10 * time.Second,
		//Addr:         ":" + os.Getenv("API_PORT"),
		//Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getConnectionString() string {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	//dbname := os.Getenv("POSTGRES_DB")
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, "test")
}

func createTableGeolocation(db *sql.DB) {
	dbname := os.Getenv("POSTGRES_DB")
	query := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname='%s'", dbname)
	stmt, err := db.Prepare(query)

	if err != nil {
		panic(err)
	}
	var result string
	err = stmt.QueryRow().Scan(&result)

	if err == sql.ErrNoRows && result == ""{
		_, err := db.Exec("CREATE DATABASE geolocation")
		if err != nil {
			panic(err)
		}

		_, er := db.Exec(`CREATE TABLE IF NOT EXISTS geolocation (
		id uuid NOT NULL,
		ipaddress TEXT NOT NULL,
		countrycode TEXT NOT NULL,
		country TEXT NOT NULL,
		city TEXT NOT NULL,
		latitude NUMERIC NOT NULL,
		longitude NUMERIC NOT NULL,
		mysteryvalue TEXT NOT NULL,
		PRIMARY KEY (id)
	)`)

		if er != nil {
			panic(er)
		}
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
