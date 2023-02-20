package postgre

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type EventNotes struct {
	Repository
}

func NewEventNotesRepository(db *sql.DB) *EventNotes {
	return &EventNotes{Repository{db}}
}
