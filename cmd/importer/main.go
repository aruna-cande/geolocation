package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/aruna-cande/geolocation/pkg/geolocation/adapters"
	"github.com/aruna-cande/geolocation/pkg/geolocation/service"
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

	dbInit := adapters.NewDBInitializer(user, password, dbname, host, port)
	db := dbInit.InitDatabase()

	defer db.Close()
	srcPath := os.Args[1]
	importGeolocationData(srcPath, db, logger)
}

func importGeolocationData(srcPath string, db *sql.DB, logger *log.Logger) {
	logger.Println("starting importer task")
	repository := adapters.NewGeolocationPostgresRepository(db)
	srv := service.NewImporterService(repository, logger)
	statistics, err := srv.ImportGeolocationData(srcPath)

	if err != nil {
		log.Println(err.Error())
	}

	log.Printf("Import duration %s. %d rows accepted and %d rows discarded", statistics.TimeElapsed, statistics.Accepted, statistics.Discarded)
}
