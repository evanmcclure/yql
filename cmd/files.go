package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func findAllDataFiles(filename, dirname, fileglob string) (result []string) {
	if len(filename) != 0 {
		fmt.Printf("Using file %s as storage.\n", filename)

		result = append(result, filename)

	} else if len(dirname) != 0 {
		fmt.Printf("Using all YAML files in dir %s as storage.\n", dirname)

		entries, err := os.ReadDir(dirname)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open dir: %v\n", err)
			os.Exit(74)
		}

		for _, ent := range entries {
			result = append(result, filepath.Join(dirname, ent.Name()))
		}

	} else if len(fileglob) != 0 {
		fmt.Printf("Using all YAML files defined by the glob %s as storage.\n", fileglob)

		matches, err := filepath.Glob(fileglob)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to use glob: %v\n", err)
			os.Exit(64)
		}

		result = append(result, matches...)
	}

	return
}
