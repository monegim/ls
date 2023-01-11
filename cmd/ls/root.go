package ls

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)
const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
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
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		s := ""
		for _, file := range files {
			fileName := file.Name()
			if !strings.HasPrefix(fileName, ".") {
				s += fmt.Sprintf("%s  ", fileName)
			}
		}
		fmt.Println(s)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI %s", err)
	}
}
