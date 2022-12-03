package main

import "testing"

func TestPuzzle1_1(t *testing.T) {
	inputsFiles := []string{"inputs/inputDay1/input1_1_1.txt", "inputs/inputDay1/input1_1_2.txt",
		"inputs/inputDay1/input1_1_3.txt",
		"inputs/inputDay1/input1_1_4.txt",
		"inputs/inputDay1/actualinput1_1.txt"}
	expectedResults := []int{24000, 21, 6, 98, 72240}
	for i := 0; i < len(inputsFiles); i++ {
		out := Puzzle1_1(inputsFiles[i])
		if out != expectedResults[i] {
			t.Errorf("Expected %d, got %d", expectedResults[i], out)
		}
	}
}

func TestPuzzle1_2(t *testing.T) {
	inputsFiles := []string{"inputs/inputDay1/input1_1_1.txt", "inputs/inputDay1/input1_1_2.txt", "inputs/inputDay1/input1_1_4.txt", "inputs/inputDay1/actualinput1_1.txt"}
	expectedResults := []int{45000, 30, 156, 210957}
	for i := 0; i < len(inputsFiles); i++ {
		out := Puzzle1_2(inputsFiles[i])
		if out != expectedResults[i] {
			t.Errorf("Expected %d, got %d", expectedResults[i], out)
		}
	}
}


