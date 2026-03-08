package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/aruna-cande/geolocation/pkg/geolocation/adapters"
	"github.com/aruna-cande/geolocation/pkg/geolocation/service"
)

const (
	maxRetries     = 5
	initialBackoff = 1 * time.Second
)

func main() {
	var logger = log.New(os.Stderr, "logger: ", log.Ldate|log.Ltime|log.Lshortfile)

	config, err := NewConfig()
	if err != nil {
		logger.Fatalf("failed to load config: %v", err)
	}

	dbInit := adapters.NewDBInitializer(config.PostgresUser, config.PostgresPassword, config.PostgresDb, config.PostgresHost, config.PostgresPort)
	db, err := connectWithRetry(dbInit, logger)
	if err != nil {
		logger.Fatalf("failed to initialize database after %d attempts: %v", maxRetries, err)
	}
	defer db.Close()

	srcPath := os.Args[1]
	importGeolocationData(srcPath, db, logger)
}

// connectWithRetry attempts to initialise the database, retrying with
// exponential backoff on failure.
func connectWithRetry(dbInit *adapters.DBInitializer, logger *log.Logger) (*sql.DB, error) {
	backoff := initialBackoff
	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		db, err := dbInit.InitDatabase()
		if err == nil {
			return db, nil
		}
		lastErr = err
		logger.Printf("database connection attempt %d/%d failed: %v — retrying in %s", attempt, maxRetries, err, backoff)
		time.Sleep(backoff)
		backoff *= 2
	}
	return nil, lastErr
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
