package postgre

import (
	"database/sql"
	errors2 "github.com/Selepok/calendar/internal/errors"
	_ "github.com/lib/pq"
	"log"
)

type UserRepository struct {
	Repository
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Repository{db}}
}

func (repo *Repository) CreateUser(login, password, timezone string) error {
	if _, err := repo.db.Exec("INSERT INTO users (login, password, timezone) VALUES ($1, $2, $3)", login, password, timezone); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetUserHashedPassword(login string) (hashedPassword string, err error) {
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

func (repo *Repository) Update(login, timezone string) error {
	if _, err := repo.db.Exec("UPDATE users SET timezone = $2 WHERE login=$1", login, timezone); err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetUserIdByLogin(login string) (id int, err error) {
	row := repo.db.QueryRow("SELECT id FROM users WHERE login=$1", login)

	err = row.Scan(&id)
	if err == sql.ErrNoRows {
		return id, errors2.NoUserFound(login)
	} else if err != nil {
		log.Println(err)
		return
	}

	return
}
