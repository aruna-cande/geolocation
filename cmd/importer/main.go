package main

import (
	"Geolocation/pkg/geolocation/adapters"
	"Geolocation/pkg/geolocation/service"
	"database/sql"
	"log"
	"os"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	var logger = log.New(os.Stderr, "logger: ", log.Ldate|log.Ltime|log.Lshortfile)
	config := NewConfig()
	user := config.PostgresUser
	password := config.PostgresPassword
	dbname := config.PostgresDb
	host := config.PostgresHost
	port := config.PostgresPort

	initDb := adapters.NewInitDb(user, password, dbname, host, port)
	db := initDb.InitDatabase()

	defer db.Close()
	srcPath := os.Args[1]
	importGeolocationData(srcPath, db, logger)
}

func importGeolocationData(srcPath string, db *sql.DB, logger *log.Logger) {
	logger.Println("starting importer task")
	repository := adapters.NewGeolocationPostgresRepository(db)
	logger.Println(repository)
	service := service.NewImporterService(repository, logger)
	logger.Println(service)
	statistics, err := service.ImportGeolocationData(srcPath)

	if err != nil {
		log.Println(err.Error())
	}

	log.Printf("Import duration %s. %d rows accepted and %d rows discarded", statistics.TimeElapsed, statistics.Accepted, statistics.Discarded)
}
