package main

import "testing"

func TestPuzzle4_1(t *testing.T) {

	inputs := []string{"inputs/inputDay4/input4_1.txt",
		"inputs/inputDay4/actualinput4_1.txt"}
	expectedResults := []int{2, 498}
	for i := 0; i < len(inputs); i++ {
		res := Puzzle4_1(inputs[i])
		if res != expectedResults[i] {
			t.Errorf("Expected %d, got %d", expectedResults[i], res)
		}
	}
}

func TestPuzzle4_2(t *testing.T) {

	inputs := []string{"inputs/inputDay4/input4_1.txt", "inputs/inputDay4/actualinput4_1.txt"}
	expectedResults := []int{4, 859}
	for i := 0; i < len(inputs); i++ {
		res := Puzzle4_2(inputs[i])
		if res != expectedResults[i] {
			t.Errorf("Expected %d, got %d", expectedResults[i], res)
		}
	}
}
