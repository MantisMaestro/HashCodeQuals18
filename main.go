package main

import (
	"fmt"
	"os"
	"time"
	"math"
)

type ride struct {
	startR       int
	startC       int
	finishR      int
	finishC      int
	earlyStart   int
	latestFinish int
	completed    bool
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
	onRide        bool
}

func main() {
	start := time.Now()

	_, err := readFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := file{}
	result := run(input)
	err = writeFile(os.Args[2], result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Main has run")

	fmt.Printf("Execute time: %v\n", time.Since(start))
}

func run(data file) []car {
	//Do the fancy stuff

	// Init cars
	var cars []car
	for i := 0; i < data.noOfCars; i++ {
		cars = append(cars, car{})
	}

	// Loop throug time 0 -> totalTime
	for t := 0; t < data.totalTime; t++ {
		// Loop through cars 0 -> noOfCars
		for _, currentCar := range cars {
			// Update position

			// Check if ride complete
			if isRideComplete(currentCar) {
				// Find next job
				nextRideIndex := findRide(currentCar, data.rides, t)
				// set job to currentRide and add to previousRides
				currentCar.currentRide = data.rides[nextRideIndex]
				currentCar.previousRides = append(currentCar.previousRides, data.rides[nextRideIndex])
				currentCar.onRide = false
				// Update rides
				data.rides[nextRideIndex].completed = true
			}
		}
	}
	return cars
}

func findRide(car car, rides []ride, currentTime int) int {
	lowestRating := 1000000
	var bestRideIndex int

	for i, currentRide := range rides{
		if currentRide.completed != true {
			rating := currentRide.earlyStart - (getDistanace(car.currentR, currentRide.startR, car.currentC, currentRide.startC) + currentTime)
			rating = int(math.Abs(float64(rating)))
			if rating < lowestRating {
				lowestRating = rating
				bestRideIndex = i
			}
		}
	}
	return bestRideIndex
}

func isRideComplete(car car) bool {
	if car.currentR == car.currentRide.finishR && car.currentC == car.currentRide.finishC{
		return true
	} else {
		return false
	}
}

func getDistance(a, b, x, y int) int {
	return int(math.Abs(float64(a-x))) + int(math.Abs(float64(b-y)))
}
