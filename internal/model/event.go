package model

import "time"

type Event struct {
	Id          int
	UserId      int
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
	Timezone    string    `json:"timezone"`
	Duration    int32     `json:"duration"`
	Notes       []string  `json:"notes"`
}

func (e *Event) OK() error {
	return nil
}
