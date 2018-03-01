package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Writes the output file
func writeFile(filename string, cars []car) (string, error) {
	fmt.Printf("Writing File: %v\n", filename)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	outputStr := ""
	for _, currentCar := range cars {
		outputStr += strconv.Itoa(len(currentCar.previousRides))
		for _, ride := range currentCar.previousRides {
			outputStr += " " + strconv.Itoa(ride.rideIndex)
		}
		outputStr += "\n"
	}
	fmt.Printf("\n\n\noutput: \n%v\n\n", outputStr)
	fmt.Fprintf(file, outputStr)

	return outputStr, nil
}
