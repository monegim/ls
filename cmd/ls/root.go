package ls

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "ls",
	Short:   "ls - ls a directory",
	Long:    "It acts like linux ls",
	Args:    cobra.ArbitraryArgs,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI %s", err)
	}
}
