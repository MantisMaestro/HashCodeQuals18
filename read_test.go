package main

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	result, err := readFile("input/a_example.in")

	var expected file
	expected.rows = 3
	expected.columns = 4
	expected.noOfCars = 2
	expected.noOfRides = 3
	expected.perRideBonus = 2
	expected.totalTime = 10

	var firstRide ride
	firstRide.startR = 2
	firstRide.startC = 0
	firstRide.finishR = 2
	firstRide.finishC = 2
	firstRide.earlyStart = 0
	firstRide.latestFinish = 9

	//Test that first line data is being read correctly
	if result.rows != expected.rows {
		t.Fatalf("Read File Error: \n%v \n%v", result.rows, expected.rows)
	}
	if result.columns != expected.columns {
		t.Fatalf("Read File Error: \n%v \n%v", result.rows, expected.rows)
	}
	if result.noOfCars != expected.noOfCars {
		t.Fatalf("Read File Error: \n%v \n%v", result.rows, expected.rows)
	}
	if result.noOfRides != expected.noOfRides {
		t.Fatalf("Read File Error: \n%v \n%v", result.rows, expected.rows)
	}
	if result.perRideBonus != expected.perRideBonus {
		t.Fatalf("Read File Error: \n%v \n%v", result.rows, expected.rows)
	}
	if result.totalTime != expected.totalTime {
		t.Fatalf("Read File Error: \n%v \n%v", result.rows, expected.rows)
	}

	//Test that rides are being read correctly
	if result.rides[2].startR != firstRide.startR {
		t.Fatalf("Read File Error: \n%v \n%v", result.rides[0].startR, firstRide.startR)
	}
	if result.rides[2].latestFinish != firstRide.latestFinish {
		t.Fatalf("Read File Error: \n%v \n%v", result.rides[0].latestFinish, firstRide.latestFinish)
	}
	
	if err != nil {
		t.Errorf("Read File Error: %v", err)
	}
}
