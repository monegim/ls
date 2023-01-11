package ls

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func isHiddenFile(name string) bool {
	return name[0] == '.'
}

func Ls(cmd *cobra.Command, args []string) {
	path := args[0]
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	s := ""
	for _, file := range files {
		fileName := file.Name()
		if  isHiddenFile(fileName) {
			continue
		} else if file.IsDir() {
			s += fmt.Sprintf(InfoColor, fileName)
			s += "  "
		} else {
			s += fmt.Sprintf("%s  ", fileName)
		}
	}
	fmt.Println(s)
}
