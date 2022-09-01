package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	isLocal := os.Getenv("IS_LOCAL")
	if isLocal == "true" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("can't load .env file!")
		}
	}

	variable := os.Getenv(key)
	if variable == "" {
		log.Fatalf(fmt.Sprintf("can't find %s", key))
	}
	return variable
}
