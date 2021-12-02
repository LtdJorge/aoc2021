package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// Storage for depth data
	var points []int

	// Read file lines to int array
	//file, err := os.Open("test-input.txt")
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		point, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		points = append(points, point)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()

	var summedPoints []int

	// Compute sliding window values
	for i := 0; i < (len(points)-2); i++ {
		newPoint := points[i] + points[i+1] + points[i+2]
		summedPoints = append(summedPoints, newPoint)
	}

	// Calculate how many increases in depth there are
	var	increases int
	for i := 0; i < len(summedPoints); i++ {
		if i != 0 {
			if summedPoints[i] > summedPoints[i-1] {
				increases++
			}
		}
	}
	fmt.Println(increases)
}