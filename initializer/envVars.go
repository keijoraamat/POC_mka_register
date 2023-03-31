package initializer

import "github.com/joho/godotenv"

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		panic("Could not load .env file")
	}
}
