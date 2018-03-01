package main

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	currentFile, err := readFile("input/a_example.in")
	if err != nil {
		panic(err)
	}
	result := run(currentFile)
	fmt.Printf("Cars: %v\n", result)
}

func TestScore(t *testing.T) {
	currentFile, err := readFile("input/a_example.in")
	if err != nil {
		panic(err)
	}
	cars := run(currentFile)
	result := getScore(cars, currentFile.perRideBonus)
	expected := 10
	if result != expected {
		t.Errorf("Score Error: %v != %v", result, expected)
	}
}
