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
}
