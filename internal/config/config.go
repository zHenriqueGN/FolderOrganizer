package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	RootFolder string 
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error on .env loading")
	}
	RootFolder = os.Getenv("ROOT_FOLDER")
}