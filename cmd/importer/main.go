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

	srcPath := os.Args[1]
	importGeolocationData(srcPath, db, logger)
}

func importGeolocationData(srcPath string, db *sql.DB, logger *log.Logger) {
	logger.Println("starting importer task")
	repository := adapters.NewGeolocationPostgresRepository(db)
	srv := service.NewImporterService(repository, logger)
	statistics, err := srv.ImportGeolocationData(srcPath)

	if err != nil {
		logger.Printf("import failed: %v", err)
	}

	logger.Printf("Import duration %s. %d rows accepted and %d rows discarded", statistics.TimeElapsed, statistics.Accepted, statistics.Discarded)
}
