package main

import (
	"fmt"
	"os"
)

// Writes the output file
func writeFile(filename string, data []car) error {
	fmt.Printf("Writing File: %v\n", filename)

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// fmt.Fprintf(file, outputStr)

	return nil
}
