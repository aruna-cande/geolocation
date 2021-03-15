package main

import (
	"Geolocation/pkg/geolocation/adapters"
	"Geolocation/pkg/geolocation/service"
	"database/sql"
	"log"
	"os"
	"strconv"
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

	defer db.Close()
	srcPath := os.Args[1]
	importGeolocationData(srcPath, db, logger)
}

func importGeolocationData(srcPath string, db *sql.DB, logger *log.Logger){
	logger.Println("starting importer task")
	repository := adapters.NewGeolocationPostgresRepository(db)
	service := service.NewImporterService(repository, logger)
	statistics, err := service.ImportGeolocationData(srcPath)

	if err != nil {
		log.Println(err.Error())
	}

	log.Printf("Import duration %s. %d rows accepted and %d rows discarded", statistics.TimeElapsed, statistics.Accepted, statistics.Discarded)
}
