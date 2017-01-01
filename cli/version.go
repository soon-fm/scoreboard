// Version CLI Sub Command

package cli

import (
	"fmt"

	"scoreboard/version"

	"github.com/spf13/cobra"
)

// Version sub command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the build version.",
	Run:   versionRunFn,
}

// Version sub command run method
func versionRunFn(cmd *cobra.Command, args []string) {
	v, err := version.Version()
	if err != nil {
		v = "unknown"
	}
	fmt.Println(fmt.Sprintf("Version: %s", v))
	bt, err := version.BuildTime()
	if err != nil {
		fmt.Println(fmt.Sprintf("Build Time: uknown"))
	} else {
		fmt.Println(fmt.Sprintf(
			"Build Time: %s",
			bt.Format("Monday Janurary 2 at 15:04:05")))
	}
}
