package main

import (
	"fmt"
	"time"
)

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
