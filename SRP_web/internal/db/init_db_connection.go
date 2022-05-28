package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	urlFormat = "postgresql://%s:%s@%s:%s/%s?sslmode=disable"
	dbConn    *sqlx.DB
)

func PingDB() (err error) {
	dbConn, err = sqlx.Open("postgres", getURL())
	if err != nil {
		return
	}
	if err = dbConn.Ping(); err != nil {
		return
	}

	return
}

func getURL() string {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgDB := os.Getenv("POSTGRES_DB")

	return fmt.Sprintf(urlFormat, pgUser, pgPassword, pgHost, pgPort, pgDB)
}
