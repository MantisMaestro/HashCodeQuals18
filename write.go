package main

import (
	"fmt"
	"os"
	"strconv"
)

// Writes the output file
func writeFile(filename string, cars []car) (string, error) {
	fmt.Printf("Writing File: %v\n", filename)

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	outputStr := ""
	for _, currentCar := range cars {
		outputStr += strconv.Itoa(len(currentCar.previousRides))
		for _, ride := range currentCar.previousRides {
			outputStr += " " + strconv.Itoa(ride.rideIndex)
		}
		outputStr += "\n"
	}
	fmt.Printf("\n\n\noutput: \n%v\n\n", outputStr)
	fmt.Fprintf(f, outputStr)

	return outputStr, nil
}
