package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func createStore(files []string) (result map[string]struct{}) {
	result = make(map[string]struct{})

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

		result[f] = file
	}

	return
}
