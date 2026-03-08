package adapters

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
)

// dbNamePattern restricts database names to safe characters only.
var dbNamePattern = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

// DBInitializer handles database creation and schema setup.
type DBInitializer struct {
	user     string
	password string
	dbname   string
	host     string
	port     int64
}

// NewDBInitializer creates a new DBInitializer with the given connection parameters.
func NewDBInitializer(user string, password string, dbname string, host string, port int64) *DBInitializer {
	return &DBInitializer{
		user:     user,
		password: password,
		dbname:   dbname,
		host:     host,
		port:     port,
	}
}

func (i DBInitializer) getConnectionString(dbname string) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d database=%s sslmode=disable", i.user, i.password, i.host, i.port, dbname)
}

// InitDatabase creates the database if it doesn't exist, ensures the schema is
// in place, and returns an open *sql.DB connection.
func (i DBInitializer) InitDatabase() (*sql.DB, error) {
	if !dbNamePattern.MatchString(i.dbname) {
		return nil, fmt.Errorf("invalid database name %q: must contain only letters, digits, and underscores", i.dbname)
	}

	postgresDB, err := sql.Open("postgres", i.getConnectionString("postgres"))
	if err != nil {
		return nil, fmt.Errorf("connecting to postgres: %w", err)
	}
	defer postgresDB.Close()

	query := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname='%s'", i.dbname)
	stmt, err := postgresDB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("preparing database check query: %w", err)
	}

	var result string
	err = stmt.QueryRow().Scan(&result)

	if err == sql.ErrNoRows && result == "" {
		query := fmt.Sprintf("CREATE DATABASE %s", i.dbname)
		_, err := postgresDB.Exec(query)
		if err != nil {
			return nil, fmt.Errorf("creating database %s: %w", i.dbname, err)
		}
	}

	connectionString := i.getConnectionString(i.dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("connecting to database %s: %w", i.dbname, err)
	}

	if err := i.createTableGeolocations(db); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func (i DBInitializer) createTableGeolocations(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS geolocations_data (
		id SERIAL PRIMARY KEY,
		ipaddress TEXT UNIQUE NOT NULL,
		countrycode TEXT NOT NULL,
		country TEXT NOT NULL,
		city TEXT NOT NULL,
		latitude NUMERIC NOT NULL,
		longitude NUMERIC NOT NULL,
		mysteryvalue TEXT NOT NULL
	)`)

	if err != nil {
		return fmt.Errorf("creating geolocations_data table: %w", err)
	}
	return nil
}
