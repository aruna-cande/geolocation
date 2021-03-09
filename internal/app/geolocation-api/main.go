package geolocation_api

import (
	"Geolocation/internal/pkg/geolocation"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/gorilla/context"
	"time"
)

func main(){
	connectionString := getConnectionString()
	db, err := sql.Open("mysql", connectionString)

	if err != nil{
		fmt.Sprintln("erro", err)
	}

	postgresRepository := geolocation.NewGeolocationFirestoreRepository(db)
	geolocationService := geolocation.NewService(postgresRepository)

	r := mux.NewRouter()

	n := negroni.New(
		negroni.NewLogger(),
	)

	http.Handle("/", r)
	CreateGeolocationHandler(r, *n, geolocationService)
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + os.Getenv("API_PORT"),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getConnectionString() string{
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
}

func createTableGeolocation(){
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS geolocation (
		id UUID NOT NULL,
		ipaddress TEXT NOT NULL,
		countrycode NOT NULL DEFAULT,
		country TEXT NOT NULL,
		city TEXT NOT NULL,
		latitude TEXT NOT NULL,
		longitude TEXT NOT NULL,
		mysteryvalue TEXT NOT NULL,
		PRIMARY KEY (username)
	)`)

	if err != nil {
		panic(err)
	}
}