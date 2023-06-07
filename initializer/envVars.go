package initializer

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if os.Getenv("REGISTER_ENV") != "" {
		err := godotenv.Load()
		if err != nil {
			panic("Could not load .env file")
		}
	}
	

}
