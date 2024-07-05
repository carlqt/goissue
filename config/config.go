package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type EnvConfig struct {
	DatabaseURL string
	DBHost      string
	DBName      string
	DBPort      string
	DBUser      string
	DBPassword  string
	Port        string
	JWTSecret   []byte
}

// NewConfig creates a new config using the environment variables.
// The variables are loaded from the .env file.
// The EnvConfig struct ensures helps with type safety and auto completion.
func NewConfig() *EnvConfig {
	return &EnvConfig{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		DBHost:      os.Getenv("DB_HOST"),
		DBName:      os.Getenv("DB_NAME"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		JWTSecret:   []byte(os.Getenv("JWT_SECRET")),
	}
}

var DBInstance *sql.DB

func NewDB() *sql.DB {
	if DBInstance != nil {
		return DBInstance
	}

	config := NewConfig()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)

	DBInstance, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = DBInstance.Ping()
	if err != nil {
		panic(err)
	}

	return DBInstance
}
