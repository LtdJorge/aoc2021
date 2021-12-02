package util

import (
	"bufio"
	"log"
	"os"
)

func ReadLinesFromFileToSlice(fileName string, slice []string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	var value [] string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		point := scanner.Text()
		value = append(value, point)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		return nil
	}

	return value
}
