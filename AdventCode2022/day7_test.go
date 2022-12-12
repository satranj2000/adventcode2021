package main

import (
	"testing"
)

func TestPuzzle7_1(t *testing.T) {
	inputs := []string{"inputs/inputDay7/input7_1.txt", "inputs/inputDay7/actualinput7_1.txt"}
	outputs := []int{95437, 0}
	for i := 0; i < len(inputs); i++ {
		out := Puzzle7_1(inputs[i])
		if out != outputs[i] {
			t.Errorf("Expected %v, got %v", outputs[i], out)
		}

	}

}
