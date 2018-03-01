package main

import (
	"bufio"
	"fmt"
	"os"
)

// Reads input file
// Params: name of the input file
// Returns: the input file object and error
func readFile(filename string) (string, error) {
	fmt.Printf("Read File: %v\n", filename)

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	// Loop through the file
	for s.Scan() {
		//row := strings.Fields(s.Text())
	}

	return "yes", nil
}
