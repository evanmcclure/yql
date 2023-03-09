package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
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

	files := findAllDataFiles(filename, dirname, fileglob)

	fmt.Println(files)

	for _, f := range files {
		data, err := os.ReadFile(f)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Skipping file %s: %v\n", f, err)
			continue
		}

		var file struct{}
		if err := yaml.Unmarshal(data, &file); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse YAML file %s: %v\n", f, err)
			continue
		}
	}

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
