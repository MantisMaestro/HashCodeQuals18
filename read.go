package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

// Reads input file
// Params: name of the input file
// Returns: the input file object
func readFile(filename string) file {
	fmt.Printf("Read File: %v\n", filename)
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	index := 0
	// var newFile file
	// Loop through the file
	for s.Scan() {
		switch index {

		case 0:

		case 1:

		case 2:

		default:

		}
		index++
	}
	return newFile
}
