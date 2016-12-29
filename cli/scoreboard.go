// Root CLI Entry Command

package cli

import (
	"scoreboard/config"
	"scoreboard/logger"
	"scoreboard/run"

	"github.com/spf13/cobra"
)

// Root Cobra CLI command
var scoreboardCmd = &cobra.Command{
	Use:              "scoreboard",
	Short:            "SOON_ FM Scoreboard.",
	Long:             "Stores the scores of users in the SOON_ FM system.",
	PersistentPreRun: scoreboardPreRunFn,
	Run:              scoreboardRunFn,
}

func init() {
	// Add CLI Flags
	scoreboardCmd.PersistentFlags().StringP(
		"config",
		"c",
		"",
		"Optional absolute path to toml config file")
	scoreboardCmd.PersistentFlags().StringP(
		"log-level",
		"l",
		"",
		"Logging log level. One of 'debug', 'info', 'warn', 'error'")
	// Add Sub Commands
	scoreboardCmd.AddCommand(versionCmd)
}

// Scoreboard CLI pre run method, loads in configuration from file
func scoreboardPreRunFn(cmd *cobra.Command, args []string) {
	// Bind flags to viper values
	config.BindLogLevelFlag(cmd.Flags().Lookup("log-level"))
	// Read in config from file
	configPath, _ := cmd.Flags().GetString("config")
	if err := config.Read(configPath); err != nil {
		logger.ConsoleOutput(true)
		logger.SetFormat("text")
		logger.WithError(err).Warn("unable to load config file")
	}
	// Setup Global Logger
	logger.Setup(config.NewLogConifg())
}

// Scoreboard CLI run method
func scoreboardRunFn(cmd *cobra.Command, args []string) {
	run.Run()
}

// Exported exec method called by main
func Exec() error {
	return scoreboardCmd.Execute()
}
