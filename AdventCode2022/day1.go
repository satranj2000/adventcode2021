package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func Readinputdata(filename string) []int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var values []int
	for scanner.Scan() {
		//println(scanner.Text())
		val, _ := strconv.Atoi(scanner.Text())
		values = append(values, val)
	}
	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()
	return values
}

func Puzzle1_1(filename string) int {

	arr := Readinputdata(filename)
	maxCal := 0
	addCal := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			if addCal > maxCal {
				maxCal = addCal
			}
			addCal = 0
		} else {
			addCal += arr[i]
		}
	}
	if addCal > maxCal {
		maxCal = addCal
	}
	return maxCal

}

func Puzzle1_2(filename string) int {
	arr := Readinputdata(filename)
	addCal := 0
	elvSnacks := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			elvSnacks = append(elvSnacks, addCal)
			addCal = 0
		} else {
			addCal += arr[i]
		}
	}
	elvSnacks = append(elvSnacks, addCal);
	
	sort.Sort(sort.Reverse(sort.IntSlice(elvSnacks)))
	return elvSnacks[0] + elvSnacks[1] + elvSnacks[2]
}
