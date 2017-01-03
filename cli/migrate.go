// CLI for handling database migrations

package cli

import (
	"scoreboard/db/migrate"
	"scoreboard/logger"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Handle DB Migrations.",
	Run:   migrateCmdFn,
}

func migrateCmdFn(cmd *cobra.Command, args []string) {
	cmd.Help()
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Run Databse Migrations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := migrate.Upgrade(); err != nil {
			logger.WithError(err).Fatal("failed to run db migrations")
		}
	},
}

var migrateRollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback Database Migrations",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		reset, _ := cmd.Flags().GetBool("all")
		to, _ := cmd.Flags().GetInt("to")
		if reset && to == 0 { // Rollback all migrations
			err = migrate.Rollback(0)
		} else if !reset && to > 0 { // Rollback to specific migration
			err = migrate.Rollback(to)
		} else if !reset && to == 0 {
			err = migrate.Downgrade() // Rollback latest migration
		}
		if err != nil {
			logger.WithError(err).Fatal("failed to rollback migrations")
		}
	},
}

func init() {
	// Bind flags
	migrateRollbackCmd.Flags().Bool("reset", false, "Reset DB migrations to 0, all data will be lost.")
	migrateRollbackCmd.Flags().Int("to", 0, "Rollback to a specific migration")
	// Add Sub Commands
	migrateCmd.AddCommand(migrateUpCmd, migrateRollbackCmd)
}
