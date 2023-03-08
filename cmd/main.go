package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var dirname string
var fileglob string
var filename string

func init() {
	flag.StringVar(&dirname, "dir", "", "name of a directory of YAML files to select over")
	flag.StringVar(&fileglob, "glob", "", "file glob of YAML to select over")
	flag.StringVar(&filename, "file", "", "name of the YAML file to select over")
}

func main() {
	parseFlags()

	var files []string
	if len(filename) != 0 {
		fmt.Printf("Using file %s as storage.\n", filename)

		files = append(files, filename)

	} else if len(dirname) != 0 {
		fmt.Printf("Using all YAML files in dir %s as storage.\n", dirname)

		entries, err := os.ReadDir(dirname)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open dir: %v\n", err)
			os.Exit(74)
		}

		for _, ent := range entries {
			files = append(files, ent.Name())
		}

	} else if len(fileglob) != 0 {
		fmt.Printf("Using all YAML files glob %s as storage.\n", dirname)

		matches, err := filepath.Glob(fileglob)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to use glob: %v\n", err)
			os.Exit(64)
		}

		files = append(files, matches...)
	}

	fmt.Println(files)

}

func parseFlags() {
	flag.Usage = func() {
		fmt.Println("Usage of yql:")
		fmt.Println("")
		fmt.Println("Use one of either -dir, -glob or -file as a storage flag.")
		fmt.Println("")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(filename) == 0 && len(dirname) == 0 && len(fileglob) == 0 {
		fmt.Fprintln(os.Stderr, "Missing a storage flag.")
		os.Exit(64)
	}
}
