// Root CLI Entry Command

package cli

import (
	"scoreboard/logger"
	"scoreboard/run"

	"github.com/spf13/cobra"
)

// Root Cobra CLI command
var scoreboardCmd = &cobra.Command{
	Use:   "scoreboard",
	Short: "SOON_ FM Scoreboard.",
	Long:  "Stores the scores of users in the SOON_ FM system.",
	Run:   scoreboardRunFn,
}

func init() {
	// --config/-c Flag
	scoreboardCmd.PersistentFlags().StringP(
		"config",
		"c",
		"",
		"Optional absolute path to toml config file")
	run.BindConfigPathFlag(scoreboardCmd.PersistentFlags().Lookup("config"))
	// --log-level/-l Flag
	scoreboardCmd.PersistentFlags().StringP(
		"log-level",
		"l",
		"",
		"Logging log level. One of 'debug', 'info', 'warn', 'error'")
	logger.BindLogLevelFlag(scoreboardCmd.PersistentFlags().Lookup("log-level"))
	// Add Sub Commands
	scoreboardCmd.AddCommand(versionCmd, importCmd)
}

// Scoreboard CLI run method
func scoreboardRunFn(cmd *cobra.Command, args []string) {
	if err := run.Run(); err != nil {
		logger.WithError(err).Fatal("runtime error")
	}
}

// Exported exec method called by main
func Exec() error {
	return scoreboardCmd.Execute()
}
