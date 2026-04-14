// package config

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// type Config struct {
// 	Port      string
// 	DBUrl     string
// 	JWTSecret string
// }

// func LoadConfig() *Config {
// 	// err := godotenv.Load("../.env")
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("Running without .env file (Docker mode)")
// 	}
// 	if err != nil {
// 		log.Println("No .env file found")
// 	}

// 	dbUrl := "postgres://" +
// 		os.Getenv("DB_USER") + ":" +
// 		os.Getenv("DB_PASSWORD") + "@" +
// 		os.Getenv("DB_HOST") + ":" +
// 		os.Getenv("DB_PORT") + "/" +
// 		os.Getenv("DB_NAME") +
// 		"?sslmode=disable"

// 	return &Config{
// 		Port:      os.Getenv("PORT"),
// 		DBUrl:     dbUrl,
// 		JWTSecret: os.Getenv("JWT_SECRET"),
// 	}
// }

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBUrl     string
	JWTSecret string
}

func LoadConfig() *Config {
	// Load .env for local development (ignored in Docker)
	if err := godotenv.Load(); err != nil {
		log.Println("Running without .env file (Docker mode)")
	}

	// Validate required env variables
	requiredVars := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"PORT",
		"JWT_SECRET",
	}

	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Missing required environment variable: %s", v)
		}
	}

	// Build DB URL
	// dbUrl := "postgres://" +
	// 	os.Getenv("DB_USER") + ":" +
	// 	os.Getenv("DB_PASSWORD") + "@" +
	// 	os.Getenv("DB_HOST") + ":" +
	// 	os.Getenv("DB_PORT") + "/" +
	// 	os.Getenv("DB_NAME") +
	// 	"?sslmode=disable"

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Println("DB_HOST not set, defaulting to localhost")
		dbHost = "localhost"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	log.Println("ENV CHECK → DB_HOST:", dbHost)

	dbUrl := "postgres://" +
		dbUser + ":" +
		dbPassword + "@" +
		dbHost + ":" +
		dbPort + "/" +
		dbName +
		"?sslmode=disable"

	log.Println("Connecting to DB at:", dbHost)

	return &Config{
		Port:      os.Getenv("PORT"),
		DBUrl:     dbUrl,
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}
