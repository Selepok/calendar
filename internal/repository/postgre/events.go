package postgre

import (
	"database/sql"
	"fmt"
	"github.com/Selepok/calendar/internal/model"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

type Events struct {
	Repository
}

func NewCalendarRepository(db *sql.DB) *Events {
	return &Events{Repository{db}}
}

func (events *Events) Create(event model.Event) (id int32, err error) {
	stmt := `INSERT INTO
				events (user_id, title, description, time, timezone, duration)
			 VALUES 
				($1, $2, $3, $4, $5, $6)
			 RETURNING id`
	if err = events.db.QueryRow(
		stmt,
		event.UserId,
		event.Title,
		event.Description,
		event.Time,
		event.Timezone,
		event.Duration,
	).Scan(&id); err != nil {
		return
	}
	var rows []string
	for _, note := range event.Notes {
		rows = append(rows, fmt.Sprintf("(%d, '%s')", id, note))
	}

	stmt = `INSERT INTO
				event_notes (event_id, item)
			 VALUES `

	if _, err = events.db.Exec(stmt + strings.Join(rows, ",")); err != nil {
		log.Println(err.Error())
	}

	return
}
