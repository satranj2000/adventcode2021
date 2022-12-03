package main

import (
	"bufio"
	"log"
	"os"
)

func ReadFiletoStrArray(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var values []string
	for scanner.Scan() {
		valStr := scanner.Text()
		values = append(values, valStr)
	}
	file.Close()
	return values
}
