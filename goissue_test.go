package goissue

import (
	"database/sql"
	"gotest_issue/config"
	integration_test "gotest_issue/test"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db := config.NewDB()
	truncateTables := func() {
		integration_test.TruncateAllTables(db)
	}

	t.Cleanup(truncateTables)

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
