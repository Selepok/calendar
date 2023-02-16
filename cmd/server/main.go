package main

import (
	"github.com/Selepok/calendar/internal/config"
	"github.com/Selepok/calendar/internal/repository/postgre"
	http2 "github.com/Selepok/calendar/internal/server/http"
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

	valid := &validator.Service{}

	userService := user.NewService(postgre.NewUserRepository(app.DB.DSN))
	server := http2.NewServer(valid, userService, app)
	router := http2.NewRouter(server)

	http.ListenAndServe(":5000", router)
}
