package main

import (
	"fmt"
	"os"

	"github.com/tmshv/fsdup/internal/scan"
)

func main() {
	rootDir := os.Args[1]
	scanner := scan.New(scan.Extensions)
	if err := scanner.Walk(rootDir); err != nil {
		fmt.Printf("error scanning directory: %v\n", err)
	}
}
