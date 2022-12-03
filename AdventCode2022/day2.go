package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Readinputdata2(filename string) [][2]string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var values [][2]string
	for scanner.Scan() {
		//println(scanner.Text())
		valStr := scanner.Text()
		valArr := strings.Split(valStr, " ")
		storVal := [2]string{valArr[0], valArr[1]}
		values = append(values, storVal)
	}
	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()
	return values
}

func Puzzle2_1(inputfile string) int {
	arr := Readinputdata2(inputfile)
	tot := 0
	rpsMap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	for i := 0; i < len(arr); i++ {
		tot += (getRockPaperScissorsScore(arr[i]) + rpsMap[arr[i][1]])

	}
	return tot
}

func getRockPaperScissorsScore(input [2]string) int {

	outcomes := map[[2]string]int{
		{"A", "X"}: 3, // rock , rock
		{"A", "Y"}: 6, // rock , paper
		{"A", "Z"}: 0, // rock , scissors
		{"B", "X"}: 0, // paper, rock
		{"B", "Y"}: 3, // paper, paper
		{"B", "Z"}: 6, // paper, scissors
		{"C", "X"}: 6, // scissors, rock
		{"C", "Y"}: 0, // scissors, paper
		{"C", "Z"}: 3, // scissors, scissors
	}

	score := outcomes[input]
	return score
}

func Puzzle2_2(inputfile string) int {
	arr := Readinputdata2(inputfile)
	tot := 0

	winOrLoseMap := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	for i := 0; i < len(arr); i++ {
		tot += (winOrLoseMap[arr[i][1]] + getResult(arr[i]))

	}
	return tot
}

func getResult(input [2]string) int {
	rpsMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	mapRst := map[[2]string]string{
		{"A", "X"}: "C", // rock wins, then return Scissors
		{"A", "Y"}: "A", // rock draws, then return rock
		{"A", "Z"}: "B", // rock loses, then return paper
		{"B", "X"}: "A", // paper wins, then return rock
		{"B", "Y"}: "B", // paper draws, then return paper
		{"B", "Z"}: "C", // paper loses, then return scissors
		{"C", "X"}: "B", // scissors wins, then return paper
		{"C", "Y"}: "C", // scissors draws, then return scissors
		{"C", "Z"}: "A", // scissors loses, then return rock
	}

	val := mapRst[input]
	return rpsMap[val]
}
