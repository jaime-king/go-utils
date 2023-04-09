package env

import (
	"github.com/joho/godotenv"
)

func Load() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		panic("Error while loading environment variables")
	}
}
