package ls

import (
	"fmt"
	"io/fs"
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

type file = os.File

func Ls(cmd *cobra.Command, args []string) {
	flags := cmd.Flags()
	a, _ := flags.GetBool("all")
	l, _ := flags.GetBool("list")
	path := ""
	if len(args) == 0 {
		path = "."
	} else {
		path = args[0]
	}
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	files, err := f.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	s := ""
	if flags.NFlag() == 0 {
		s = simpleLS(files)
	} else if l {
		s = lsL(files)
	} else if a {
		s = lsA(files)
	}
	fmt.Println(s)
}

func simpleLS(files []fs.FileInfo) string {
	s := ""
	for _, file := range files {
		fileName := file.Name()
		if isHiddenFile(fileName) {
			continue
		} else if file.IsDir() {
			s += fmt.Sprintf(InfoColor, fileName)
			s += "  "
		} else {
			s += fmt.Sprintf("%s  ", fileName)
		}
	}
	return s
}

func lsA(files []fs.FileInfo) string {
	// ls -a
	s := ""
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() {
			s += fmt.Sprintf(InfoColor, fileName)
			s += "  "
		} else {
			s += fmt.Sprintf("%s  ", fileName)
		}
	}
	return s
}

func lsL(files []fs.FileInfo) string {
	// ls -a
	s := ""
	for _, file := range files {
		// fileName := file.Name()
		// if file.IsDir() {
		// 	s += fmt.Sprintf(InfoColor, fileName)
		// 	s += "  "
		// } else {
		// 	s += fmt.Sprintf("%s  ", fileName)
		// }
		s += fmt.Sprintf("%s %d %s %s\n", 
			file.Mode(), 

			file.Size(),
			file.ModTime(),
			file.Name(),
				)	
	}
	return s
}
