package config

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"log"
)

// Application holds application configuration values
type Application struct {
	SecretKey               string `env:"SECRET_KEY"`
	TokenExpirationDuration int64  `env:"TOKEN_EXPIRATION_TIME_IN_MINUTES"`
	DB                      *Database
}

type Database struct {
	DSN string `env:"DSN"`
}

func GetApplication() Application {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	ctx := context.TODO()
	var app Application

	if err := envconfig.Process(ctx, &app); err != nil {
		log.Println(err)
	}

	return app
}
