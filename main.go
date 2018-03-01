package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

type ride struct {
	rideIndex       int
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
	_, err = writeFile(os.Args[2], result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Score: %v\n", getScore(result, input.perRideBonus))
	fmt.Printf("Main has run\n")

	fmt.Printf("Execute time: %v\n", time.Since(start))
}

func run(data file) []car {
	//Do the fancy stuff
	fmt.Printf("File: %v\n", data)
	// Init cars
	var cars []car
	for i := 0; i < data.noOfCars; i++ {
		cars = append(cars, car{})
	}

	// Init rides
	for i, currentCar := range cars {
		nextRideIndex := findRide(currentCar, data.rides, 0)
		fmt.Printf("nextRideIndex: %v\n", nextRideIndex)
		// set job to currentRide and add to previousRides
		currentCar.currentRide = data.rides[nextRideIndex]
		nextRide := data.rides[nextRideIndex]
		nextRide.completed = true
		currentCar.previousRides = append(currentCar.previousRides, nextRide)
		currentCar.onRide = false
		// Update rides
		data.rides[nextRideIndex].completed = true

		cars[i] = currentCar
		fmt.Printf("Init car: %v\n", currentCar)
	}

	// Loop throug time 0 -> totalTime
	for t := 0; t < data.totalTime; t++ {
		fmt.Printf("Time: %v\n", t)
		// Loop through cars 0 -> noOfCars
		for i, currentCar := range cars {
			// Update position
			currentCar = updatePosition(currentCar)

			// Wait for Ride
			if !currentCar.onRide {
				if currentCar.currentC == currentCar.currentRide.startC &&
					currentCar.currentR == currentCar.currentRide.startR {
					if t >= currentCar.currentRide.earlyStart {
						if len(currentCar.previousRides) > 0 {
							currentCar.previousRides[len(currentCar.previousRides)-1].startTime = t
							currentCar.onRide = true
						}
					}
				}
			}

			// Check if ride complete
			if isRideComplete(currentCar) && currentCar.onRide {
				if t <= currentCar.currentRide.latestFinish {
					if len(currentCar.previousRides) > 0 {
						currentCar.previousRides[len(currentCar.previousRides)-1].completedOnTime = true
					}
				}
				// Find next job
				nextRideIndex := findRide(currentCar, data.rides, t)
				// set job to currentRide and add to previousRides
				currentCar.currentRide = data.rides[nextRideIndex]
				nextRide := data.rides[nextRideIndex]
				nextRide.completed = true
				currentCar.previousRides = append(currentCar.previousRides, nextRide)
				currentCar.onRide = false
				// Update rides
				data.rides[nextRideIndex].completed = true
			}

			cars[i] = currentCar

			fmt.Printf("\tCar %v:\n", i)
			fmt.Printf("\t\tPosition: r = %v c = %v\n", currentCar.currentR, currentCar.currentC)
			fmt.Printf("\t\tonRide: %v \n", currentCar.onRide)
			fmt.Printf("\t\tCurrent Ride: %v\n", currentCar.currentRide)
			fmt.Printf("\t\tPrevious Rides: %v\n", currentCar.previousRides)
		}
		fmt.Printf("\nRides: %v\n\n", data.rides)
	}
	return cars
}

func findRide(car car, rides []ride, currentTime int) int {
	lowestRating := 1000000
	var bestRideIndex int

	for i, currentRide := range rides {
		if currentRide.completed != true {
			rating := currentRide.earlyStart - (getDistance(car.currentR, currentRide.startR, car.currentC, currentRide.startC) + currentTime)
			rating = int(math.Abs(float64(rating)))
			if rating < lowestRating {
				lowestRating = rating
				bestRideIndex = i
			}
		}
	}
	fmt.Printf("FindRide: index: %v ride: %v\n", bestRideIndex, rides[bestRideIndex])
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
	fmt.Printf("\t\t\tUpdating Postion\n")
	// Check if heading to finsh, or start of current ride
	if car.onRide {
		dR := car.currentR - car.currentRide.finishR
		dC := car.currentC - car.currentRide.finishC
		fmt.Printf("\t\t\t\tdR: %v dC: %v\n", dR, dC)
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
		fmt.Printf("\t\t\t\tdR: %v dC: %v\n", dR, dC)
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
				distance := getDistance(currentRide.startR, currentRide.startC, currentRide.finishR, currentRide.finishC)
				score += distance
			}
			if currentRide.startTime == currentRide.earlyStart {
				score += bonus
			}
		}
	}
	return score
}
