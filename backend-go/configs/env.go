package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
    // Try to load .env.test first, then fall back to .env
    err := godotenv.Load(".env.test")
    if err != nil {
        err = godotenv.Load()
    }
    
    if err != nil {
        // For testing, return a default local MongoDB URI
        return "mongodb://localhost:27017"
    }

    uri := os.Getenv("MONGOURI")
    if uri == "" {
        log.Fatal("MONGOURI not set in the .env or .env.test file")
    }
    return uri
}
