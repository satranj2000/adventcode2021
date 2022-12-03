package main

import "testing"

func TestPuzzle2_1(t *testing.T) {
	inputsFiles := []string{"inputs/inputDay2/input2_1_1.txt", "inputs/inputDay2/actualinput2_1.txt"}
	expectedResults := []int{15, 11873}
	for i := 0; i < len(inputsFiles); i++ {
		out := Puzzle2_1(inputsFiles[i])
		if out != expectedResults[i] {
			t.Errorf("Expected %d, got %d", expectedResults[i], out)
		}
	}
}
func TestPuzzle2_2(t *testing.T) {
	inputsFiles := []string{"inputs/inputDay2/input2_1_1.txt", "inputs/inputDay2/actualinput2_1.txt"}
	expectedResults := []int{12, 12014}
	for i := 0; i < len(inputsFiles); i++ {
		out := Puzzle2_2(inputsFiles[i])
		if out != expectedResults[i] {
			t.Errorf("Expected %d, got %d", expectedResults[i], out)
		}
	}
}
