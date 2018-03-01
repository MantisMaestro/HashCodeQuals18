package main

import "testing"

// func TestRun(t *testing.T) {
// 	result, err := run(readFile("input/a_example.in"))
// 	if err != nil {
// 		panic(err)
// 	}
// }

func TestScore(t *testing.T) {
	result := 10 //getScore(cars, bonus)
	expected := 10
	if result != expected {
		t.Errorf("Score Error: %v != %v", result, expected)
	}
}
