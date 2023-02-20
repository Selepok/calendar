package postgre

import (
	"database/sql"
	"log"
)

type Repository struct {
	db *sql.DB
}

func InitPostgresConnection(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
