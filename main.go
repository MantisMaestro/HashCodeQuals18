package main

import (
	"fmt"
	"time"
)

type ride struct {
	startR       int
	startC       int
	finishR      int
	finishC      int
	earlyStart   int
	latestFinish int
}

type file struct {
	rows         int
	columns      int
	noOfCars     int
	noOfRides    int
	rides        []ride
	perRideBonus int
	totalTime    int
}

type car struct {
	currentR      int
	currentC      int
	currentRide   ride
	previousRides []ride
}

func main() {
	start := time.Now()

	readFile("Hello World")
	writeFile("Goodbye World")
	fmt.Printf("Main has run")

	fmt.Printf("Execute time: %v\n", time.Since(start))
}

func run() {
	//Do the fancy stuff

	// Loop throug time 0 -> totalTime
	// Loop through cars 0 -> noOfCars
	// Check if car NOT on job
	// Find next job
	// set job to currentRide and add to previousRides
}
