package integration_test

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand/v2"

	"golang.org/x/crypto/bcrypt"
)

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
