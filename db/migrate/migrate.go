// Simple package for handling database migrations

package migrate

import (
	"database/sql"

	"scoreboard/db"
	"scoreboard/logger"
	"scoreboard/run"

	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
)

var log = logger.WithField("pkg", "db/migrate")

// Run function with DB context
func Migrate(fn func(*gomigrate.Migrator) error) error {
	if err := run.ReadConfig(); err != nil {
		return err
	}
	config := db.NewConfig()
	dbconn, err := sql.Open("postgres", config.ConnectionURL())
	if err != nil {
		return err
	}
	defer dbconn.Close()
	migrator, err := gomigrate.NewMigrator(
		dbconn,
		gomigrate.Postgres{},
		config.MigrationPath())
	if err != nil {
		return err
	}
	return fn(migrator)
}

// Run upgrade database migrations
func Upgrade() error {
	return Migrate(func(migrator *gomigrate.Migrator) error {
		return migrator.Migrate()
	})
}

// Run downgrade database migrations
func Downgrade() error {
	return Migrate(func(migrator *gomigrate.Migrator) error {
		return migrator.Rollback()
	})
}

// Rollback migrations to n number, if 0 then all migrations are rolled back
// 0 == all migrations
// >0 == rollback to n migrations
func Rollback(n int) error {
	return Migrate(func(migrator *gomigrate.Migrator) error {
		if n > 0 {
			return migrator.RollbackN(n)
		} else {
			return migrator.RollbackAll()
		}
		return nil
	})
}
