package main

import "testing"

func TestPuzzle6_1(t *testing.T) {
	inputs := []string{"inputs/inputDay6/input6_1.txt", "inputs/inputDay6/actualinput6_1.txt"}
	//, "inputs/inputDay5/actualinput5_1.txt"} //"inputs/inputDay4/actualinput4_1.txt"
	outputs := [][]int{{7, 5, 6, 10, 11}, {1909}}

	for i := 0; i < len(inputs); i++ {
		out := Puzzle6_1(inputs[i])
		for j := 0; j < len(outputs[i]); j++ {
			if out[j] != outputs[i][j] {
				t.Errorf("Expected %v, got %v", out, outputs[i])
			}
		}
	}

}

func TestPuzzle6_2(t *testing.T) {
	inputs := []string{"inputs/inputDay6/input6_1.txt", "inputs/inputDay6/actualinput6_1.txt"}
	//, "inputs/inputDay5/actualinput5_1.txt"} //"inputs/inputDay4/actualinput4_1.txt"
	outputs := [][]int{{19, 23, 23, 29, 26}, {1909}}

	for i := 0; i < len(inputs); i++ {
		out := Puzzle6_2(inputs[i])
		for j := 0; j < len(outputs[i]); j++ {
			if out[j] != outputs[i][j] {
				t.Errorf("Expected %v, got %v", out, outputs[i])
			}
		}
	}

}
