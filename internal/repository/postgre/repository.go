package postgre

import (
	"database/sql"
	"fmt"
	errors2 "github.com/Selepok/calendar/internal/errors"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(dsn string) *Repository {
	// create connection
	dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_URL"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Repository{db}
}

func (repo Repository) CreateUser(login, password, timezone string) error {
	if _, err := repo.db.Exec("INSERT INTO users VALUES ($1, $2, $3)", login, password, timezone); err != nil {
		return err
	}
	return nil
}

func (repo Repository) GetUserHashedPassword(login string) (hashedPassword string, err error) {
	row := repo.db.QueryRow("SELECT password FROM users WHERE login=$1", login)

	err = row.Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return "", errors2.NoUserFound(login)
	} else if err != nil {
		log.Println(err)
		return
	}

	return hashedPassword, nil
}
