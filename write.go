package main

import (
	"fmt"
	"os"
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
		outputStr += string(len(currentCar.previousRides))
		for _, ride := range currentCar.previousRides {
			outputStr += " " + string(ride.rideIndex)
		}
		outputStr += "\n"
	}

	fmt.Fprintf(f, outputStr)

	return outputStr, nil
}
