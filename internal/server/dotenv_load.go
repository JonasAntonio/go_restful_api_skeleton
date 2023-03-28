package server

import (
	"log"

	env "github.com/joho/godotenv"
)

func init() {
	err := env.Load("../../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
