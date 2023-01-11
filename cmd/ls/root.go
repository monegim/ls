package ls

import (
	"fmt"
	"os"

	"github.com/monegim/ls/pkg/ls"
	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "ls",
	Short:   "ls - ls a directory",
	Long:    "It acts like linux ls",
	Args:    cobra.ArbitraryArgs,
	Version: version,
	Run:     ls.Ls,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI %s", err)
	}
}
