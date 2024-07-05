package goissue

import (
	"database/sql"
	integration_test "gotest_issue/test"
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
