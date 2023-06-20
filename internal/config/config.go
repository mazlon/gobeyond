//This package contains functions to load configuration from various sources such as
//  environment variables, configuration files, and command-line arguments.
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Returns ENV variable from .env file or os
func GetTheEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	env := os.Getenv(key)
	if env != "" {
		return env
	} else {
		log.Printf("Couldn't find the ENV key: %s", key)
		return ""
	}
}
