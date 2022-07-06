package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"go_get_struct/generator"
	"go_get_struct/node"
)

func main() {
	flag.Parse()
	for _, pathFile := range flag.Args() {
		absPath, err := filepath.Abs(pathFile)
		if err != nil {
			log.Fatalf("Unable to get absolute path to file %v. Error: %v", pathFile, err)
			return
		}
		file, err := os.Open(absPath)
		if err != nil {
			log.Fatalf("Failed to open file %v. Error: %v", pathFile, err)
			return
		}
		defer file.Close()
		n, err := node.Parse(file)
		if err != nil {
			log.Fatalf("Failed to parse file %v. Error: %v", pathFile, err)
			return
		}
		if err := generator.CreateStruct(n, absPath); err != nil {
			log.Fatalf("Failed to get go structure for file %v. Error: %v", pathFile, err)
			return
		}
	}

}
