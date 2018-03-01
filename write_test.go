package main

import (
	"testing"
)

func TestWriteFile(t *testing.T) {

	currentFile, err := readFile("input/a_example.in")
	if err != nil {
		panic(err)
	}
	cars := run(currentFile)

	result, err := writeFile("output/a_example.out", cars)
	if err != nil {
		panic(err)
	}
	expected := "1 0\n2 2 1"
	if result != expected {
		t.Errorf("Write File Error: Expected %v\n%v\n", expected, result)
	}
}
