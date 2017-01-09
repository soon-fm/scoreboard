// Version CLI Sub Command

package cli

import (
	"fmt"

	"scoreboard/history"
	"scoreboard/logger"
	"scoreboard/run"

	"github.com/spf13/cobra"
)

// Version sub command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import playlist history scores",
	Run: func(cmd *cobra.Command, args []string) {
		if err := run.ReadConfig(); err != nil {
			logger.WithError(err).Warn("error loading configuration")
		}
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			logger.WithError(err).Fatal("failed to import history")
		}
		if err := history.Import(path); err != nil {
			logger.WithError(err).Fatal("failed to import history")
		}
		fmt.Println("History imported")
	},
}

func init() {
	importCmd.Flags().String("path", "", "Path to json dump of playlist history")
}
