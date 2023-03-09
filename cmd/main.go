package main

import (
	"flag"
	"fmt"
	"os"
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

	store := createStore(files)

	fmt.Println(len(store))
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
