package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// Reads input file
// Params: name of the input file
// Returns: the input file object and error
func readFile(filename string) (file, error) {
	fmt.Printf("Read File: %v\n", filename)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	index := 0
	var newFile file

	// Loop through the file
	for s.Scan() {
		row := strings.Fields(s.Text())
		switch index {
		case 0:
			newFile.rows, _ = strconv.Atoi(row[0])
			newFile.columns, _ = strconv.Atoi(row[1])
			newFile.noOfCars, _ = strconv.Atoi(row[2])
			newFile.noOfRides, _ = strconv.Atoi(row[3])
			newFile.perRideBonus, _ = strconv.Atoi(row[4])
			newFile.totalTime, _ = strconv.Atoi(row[5])
			newFile.rides = make([]ride, newFile.noOfRides)
		default:
			var newRide ride
			newRide.startR, _ = strconv.Atoi(row[0])
			newRide.startC, _ = strconv.Atoi(row[1])
			newRide.finishR, _ = strconv.Atoi(row[2])
			newRide.finishC, _ = strconv.Atoi(row[3])
			newRide.earlyStart, _ = strconv.Atoi(row[4])
			newRide.latestFinish, _ = strconv.Atoi(row[5])
			newRide.rideIndex = index-1
			newFile.rides[index-1] = newRide
		}
		index++
	}
	return newFile, nil
}
