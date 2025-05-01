package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("config/.env.development")
	if err != nil {
		log.Fatal("Env variables wasnt load successfully")
	}
}