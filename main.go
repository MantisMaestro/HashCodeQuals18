package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

type ride struct {
	startR          int
	startC          int
	finishR         int
	finishC         int
	earlyStart      int
	latestFinish    int
	completed       bool
	startTime       int
	completedOnTime bool
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
	fmt.Printf("Main started\n")

	input, err := readFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	result := run(input)
	err = writeFile(os.Args[2], result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Score: %v\n", getScore(result, input.perRideBonus))
	fmt.Printf("Main has run\n")

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
			currentCar = updatePosition(currentCar)

			// Wait for Ride
			if !currentCar.onRide {
				if currentCar.currentC == currentCar.currentRide.startC &&
					currentCar.currentR == currentCar.currentRide.startR {
					if t >= currentCar.currentRide.earlyStart {
						currentCar.previousRides[len(currentCar.previousRides)-1].startTime = t
						currentCar.onRide = true
					}
				}
			}

			// Check if ride complete
			if isRideComplete(currentCar) && currentCar.onRide {
				if t <= currentCar.currentRide.latestFinish {
					currentCar.previousRides[len(currentCar.previousRides)-1].completedOnTime = true
				}
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

	for i, currentRide := range rides {
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
	if car.currentR == car.currentRide.finishR && car.currentC == car.currentRide.finishC {
		return true
	} else {
		return false
	}
}

func getDistance(a, b, x, y int) int {
	return int(math.Abs(float64(a-x))) + int(math.Abs(float64(b-y)))
}

func updatePosition(car car) car {
	// Check if heading to finsh, or start of current ride
	if car.onRide {
		dR := car.currentR - car.currentRide.finishR
		dC := car.currentC - car.currentRide.finishC
		if dR != 0 {
			if dR > 0 {
				car.currentR--
			} else {
				car.currentR++
			}
		} else if dC != 0 {
			if dC > 0 {
				car.currentC--
			} else {
				car.currentC++
			}
		}
	} else {
		dR := car.currentR - car.currentRide.startR
		dC := car.currentC - car.currentRide.startC
		if dR != 0 {
			if dR > 0 {
				car.currentR--
			} else {
				car.currentR++
			}
		} else if dC != 0 {
			if dC > 0 {
				car.currentC--
			} else {
				car.currentC++
			}
		}
	}
	return car
}

func getScore(cars []car, bonus int) int {
	var score = 0
	for _, currentCar := range cars {
		for _, currentRide := range currentCar.previousRides {
			if currentRide.completedOnTime {
				distance := 1
				score += distance
			}
			if currentRide.startTime == currentRide.earlyStart {
				score += bonus
			}
		}
	}
	return score
}
