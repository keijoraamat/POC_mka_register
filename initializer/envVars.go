package initializer

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if os.Getenv("APP_ENV") != "" {
		err := godotenv.Load()
		if err != nil {
			panic("Could not load .env file")
		}
	}

	err := os.Setenv("PORT", "3001")
	if err != nil {
		panic("could not set app port")
	}

}
