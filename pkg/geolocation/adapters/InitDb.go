package adapters

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type InitDb struct {
	user     string
	password string
	dbname   string
}

func NewInitDb(user string, password string, dbname string) *InitDb {
	return &InitDb{
		user:     user,
		password: password,
		dbname:   dbname,
	}
}

func getConnectionString(user string, password string, dbname string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
}

func (i InitDb) InitDatabase() (postgresDB *sql.DB) {
	postgresDB, err := sql.Open("postgres", getConnectionString(i.user, i.password, "postgres"))
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
			panic(err)
		}
	}

	connectionString := getConnectionString(i.user, i.password, i.dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	i.createTableGeolocations(db)

	return db
}

func (i InitDb) createTableGeolocations(db *sql.DB) {
	_, er := db.Exec(`
		CREATE TABLE IF NOT EXISTS geolocations_data (
		id SERIAL PRIMARY KEY,
		ipaddress TEXT NOT NULL,
		countrycode TEXT NOT NULL,
		country TEXT NOT NULL,
		city TEXT NOT NULL,
		latitude NUMERIC NOT NULL,
		longitude NUMERIC NOT NULL,
		mysteryvalue TEXT NOT NULL
	)`)

	if er != nil {
		panic(er)
	}
}
