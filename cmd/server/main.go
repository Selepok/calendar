package main

import (
	"github.com/Selepok/calendar/internal/repository/postgre"
	http2 "github.com/Selepok/calendar/internal/server/http"
	"github.com/Selepok/calendar/internal/services/calendar"
	"github.com/Selepok/calendar/internal/services/validator"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// main server code
	//
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	valid := &validator.Service{}
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	//TODO: create repository
	//TODO: create calendarService. Inject repository
	calendarService := calendar.NewService(postgre.NewRepository(os.Getenv("DSN")))
	server := http2.NewServer(valid, calendarService)
	//
	router := http2.NewRouter(server)

	http.ListenAndServe(":5000", router)
}
