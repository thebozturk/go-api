package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	// read .env file and load into environment
	err := godotenv.Load()

	// if error loading .env file, log error and exit
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	// get mongo uri from environment
	mongoIRU := os.Getenv("MONGO_URI")
	return mongoIRU
}

func EnvPort() string {
	// read .env file and load into environment
	err := godotenv.Load()

	// if error loading .env file, log error and exit
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	// get port from environment
	port := os.Getenv("PORT")
	return port
}
