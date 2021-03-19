package adapters

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type InitDb struct {
	user     string
	password string
	dbname   string
	host     string
	port     int64
}

func NewInitDb(user string, password string, dbname string, host string, port int64) *InitDb {
	return &InitDb{
		user:     user,
		password: password,
		dbname:   dbname,
		host:     host,
		port:     port,
	}
}

func (i InitDb) getConnectionString(dbname string) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d database=%s sslmode=disable", i.user, i.password, i.host, i.port, dbname)
}

func (i InitDb) InitDatabase() (postgresDB *sql.DB) {
	postgresDB, err := sql.Open("postgres", i.getConnectionString("postgres"))
	defer postgresDB.Close()

	query := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname='%s'", i.dbname)
	stmt, err := postgresDB.Prepare(query)

	if err != nil {
		panic(err)
	}
	var result string
	err = stmt.QueryRow().Scan(&result)

	if err == sql.ErrNoRows && result == "" {
		query := fmt.Sprintf("CREATE DATABASE %s", i.dbname)
		_, err := postgresDB.Exec(query)
		if err != nil {
			log.Println("unable to create database " + i.dbname)
			panic(err)
		}
	}

	connectionString := i.getConnectionString(i.dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	i.createTableGeolocations(db)

	return db
}

func (i InitDb) createTableGeolocations(db *sql.DB) {
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
		log.Println("unable to create table geolocations_data")
		panic(err)
	}
}
