package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func calculateWrappingPaperLength(l, w, h int) int {
	// Calculate surface area
	area := 2*l*w + 2*w*h + 2*h*l
	// Find the area of the smallest side
	minSide := min(min(l*w, w*h), h*l)
	// Total wrapping paper needed, including slack
	total := area + minSide
	return total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Read input from file
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Split input into lines
	lines := strings.Split(string(data), "\n")

	totalSquareFeet := 0

	// Calculate total wrapping paper needed
	for _, line := range lines {
		if line == "" {
			continue
		}
		// Split line into dimensions
		dimensions := strings.Split(line, "x")
		if len(dimensions) != 3 {
			log.Fatalf("Invalid input format: %s", line)
		}
		// Convert dimensions to integers
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])
		// Calculate wrapping paper needed for this present
		totalSquareFeet += calculateWrappingPaperLength(l, w, h)
	}

	fmt.Println("Total square feet of wrapping paper to order:", totalSquareFeet)
}
