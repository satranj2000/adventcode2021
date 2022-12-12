package main

import "testing"

func TestPuzzle8_1(t *testing.T) {

	inputs := []string{"inputs/inputDay8/input8_1.txt",
		"inputs/inputDay8/actualinput8_1.txt",
		"inputs/inputDay8/input8_2.txt",
	}

	outputs := []int{21, 1870, 1669}
	for i := 0; i < len(inputs); i++ {
		out := Puzzle8_1(inputs[i])
		if out != outputs[i] {
			t.Errorf("Expected %v, got %v", outputs[i], out)
		}
	}
}

func TestPuzzle8_2(t *testing.T) {

	inputs := []string{"inputs/inputDay8/input8_1.txt",
		"inputs/inputDay8/actualinput8_1.txt",
		"inputs/inputDay8/input8_2.txt",
	}

	outputs := []int{8, 517440, 331344}
	for i := 0; i < len(inputs); i++ {
		out := Puzzle8_2(inputs[i])
		if out != outputs[i] {
			t.Errorf("Expected %v, got %v", outputs[i], out)
		}
	}
}
