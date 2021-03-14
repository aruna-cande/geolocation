package main

import (
	"Geolocation/pkg/geolocation/adapters"
	"Geolocation/pkg/geolocation/service"
	"database/sql"
	"log"
	"os"
)

func main() {

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	initDb := adapters.NewInitDb(user, password, dbname)
	db := initDb.InitDatabase()
	defer db.Close()
	srcPath := os.Args[1]
	importGeolocationData(srcPath, db)
}

func importGeolocationData(srcPath string, db *sql.DB){
	var logger = log.New(os.Stderr, "logger: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("starting importer service")
	repository := adapters.NewGeolocationPostgresRepository(db)
	service := service.NewImporterService(repository, logger)
	statistics, err := service.ImportGeolocationData(srcPath)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Import duration %s. %d rows accepted and %d rows discarded", statistics.TimeElapsed, statistics.Accepted, statistics.Discarded)
}
