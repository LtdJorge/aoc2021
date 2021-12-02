package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	instruction string
	value int
}

func main() {
	var horizontal, depth int = 0, 0
	forward := "forward"
	down := "down"
	up := "up"

	//var pairArray = fileToSlice("test-input.txt")
	var pairArray = fileToSlice("input.txt")
	var pairs []pair

	for i := 0; i < len(pairArray); i++ {
		pair := parsePair(pairArray[i])
		pairs = append(pairs, pair)
		switch pair.instruction {
		case forward:
			horizontal += pair.value
		case up:
			depth -= pair.value
		case down:
			depth += pair.value
		default:
			log.Fatalf("Unable to parse pair: %s %d", pair.instruction, pair.value)
		}
	}
	fmt.Printf("Horizontal: %d\n", horizontal)
	fmt.Printf("Depth: %d\n", depth)
	fmt.Printf("Multiplied: %d\n", horizontal*depth)
}

func fileToSlice(filename string) []string{
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	var points []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		point := scanner.Text()
		points = append(points, point)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		return nil
	}
	return points
}

func parsePair(str string) pair {
	var parts = strings.Split(str, " ")
	var pair pair
	pair.instruction = parts[0]
	var err error
	pair.value, err = strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	return pair
}