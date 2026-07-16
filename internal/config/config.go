package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Check if .env exists
	if _, err := os.Stat(".env"); err == nil {
		// File exists → load it
		if err := godotenv.Load(); err != nil {
			log.Printf("⚠️ Could not load .env: %v", err)
		} else {
			log.Println("✅ .env file loaded")
		}
	} else if !os.IsNotExist(err) {
		// Some other filesystem error
		log.Printf("⚠️ Error checking .env: %v", err)
	} else {
		// File not found → skip silently
		log.Println("ℹ️ No .env file found, skipping")
	}
}
