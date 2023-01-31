package main

import (
	"github.com/Selepok/calendar/internal/repository/postgre"
	myServer "github.com/Selepok/calendar/internal/server/http"
	"github.com/Selepok/calendar/internal/services/calendar"
	"github.com/Selepok/calendar/internal/services/validator"
	"net/http"
	"os"
)

func main() {
	// main server code

	valid := &validator.Service{}

	//TODO: create repository
	//TODO: create calendarService. Inject repository
	calendarService := calendar.NewService(postgre.NewRepository(os.Getenv("DSN")))

	server := myServer.NewServer(valid, calendarService)

	router := myServer.NewRouter(server)

	http.ListenAndServe(":5000", router)
}
