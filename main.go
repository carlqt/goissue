package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"

	_ "github.com/lib/pq"

	"gotest_issue/repository"

	"math/rand/v2"

	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type App struct {
	Repositories *repository.Repository
	Config       *EnvConfig
	DB           *sql.DB
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func NewApp() *App {
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	config := NewConfig()
	db := newDB(config)
	repositories := repository.NewRepository(db)

	return &App{Config: config, DB: db, Repositories: repositories}
}

func newDB(config *EnvConfig) *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("cannot generate hash from string")
	}

	return string(hash), nil
}

func CreateUser(db *sql.DB) {
	username := fmt.Sprintf("fake_user+%d", rand.IntN(100))
	fakePassword := "password"

	HashPassword(fakePassword)

	sql := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	_, err := db.Query(sql, username, fakePassword)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Starting")
}
