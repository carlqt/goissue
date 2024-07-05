package repository

import (
	"database/sql"
	"fmt"
	integration_test "gotest_issue/test"
	"os"
	"testing"

	"log"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	config := NewConfig()
	db := newDB(config)
	truncateTables := func() {
		TruncateAllTables(db)
	}

	t.Run("create 5 users", func(t *testing.T) {
		t.Cleanup(truncateTables)

		expectedCount := 5

		for i := 0; i < expectedCount; i++ {
			integration_test.CreateUser(db)
		}

		actualCount := countUsers(db)

		assert.Equal(t, expectedCount, actualCount)
	})

	t.Run("create 5 users", func(t *testing.T) {
		t.Cleanup(truncateTables)
		expectedCount := 5

		for i := 0; i < expectedCount; i++ {
			integration_test.CreateUser(db)
		}

		actualCount := countUsers(db)

		assert.Equal(t, expectedCount, actualCount)
	})

	t.Run("create 5 users", func(t *testing.T) {
		t.Cleanup(truncateTables)
		expectedCount := 5

		for i := 0; i < expectedCount; i++ {
			integration_test.CreateUser(db)
		}

		actualCount := countUsers(db)

		assert.Equal(t, expectedCount, actualCount)
	})
}

func countUsers(db *sql.DB) int {
	count := 0

	db.QueryRow("SELECT count(*) FROM users").Scan(&count)

	return count
}

func TruncateAllTables(db *sql.DB) {
	query := `
		DO $$ DECLARE
			r RECORD;
		BEGIN
			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname =current_schema()) LOOP
				EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' CASCADE';
			END LOOP;
		END $$;
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
	}
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
