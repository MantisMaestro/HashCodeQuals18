package main

import (
	"os"
	"log"
	"strconv"
	"fmt"
)

func writeFile(filename string, /*data*/) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	// Init values

	// Loop though data
	for i := 0; i < len(); i++ {
		for j := 0; j < len(); j++ {
			// if cell has backbone, inc backboneTotal and add coords to list

			// if cell has router, inc routerTotal and add coords to list

		}
	}
	
	outputStr :=
	outputStr +=
	outputStr +=
	outputStr +=
	fmt.Fprintf(file, outputStr)
}
