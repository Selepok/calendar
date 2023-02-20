package main

import (
	"github.com/Selepok/calendar/internal/config"
	"github.com/Selepok/calendar/internal/repository/postgre"
	http2 "github.com/Selepok/calendar/internal/server/http"
	"github.com/Selepok/calendar/internal/services/calendar"
	"github.com/Selepok/calendar/internal/services/user"
	"github.com/Selepok/calendar/internal/services/validator"
	"log"
	"net/http"
	"os"
)

func main() {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)
	app := config.GetApplication()
	dbConn := postgre.InitPostgresConnection(app.DB.DSN)

	valid := &validator.Service{}
	userRepository := postgre.NewUserRepository(dbConn)
	userService := user.NewService(userRepository)
	calendarService := calendar.NewEventsService(postgre.NewCalendarRepository(dbConn), userRepository)
	server := http2.NewServer(valid, userService, calendarService, app)
	router := http2.NewRouter(server)

	http.ListenAndServe(":5000", router)
}
