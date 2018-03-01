package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	result, err := readFile("main.go")
	if err != nil {
		t.Errorf("Read File Error: %v", err)
	}
	expected := "yes"
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Read File Error: %v \t%v\n", result, expected)
	}
}
