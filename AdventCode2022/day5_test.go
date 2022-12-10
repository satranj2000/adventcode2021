package main

import (
	"testing"
)

func TestPuzzle5_1(t *testing.T) {
	inputs := []string{"inputs/inputDay5/input5_1.txt", "inputs/inputDay5/actualinput5_1.txt"} //"inputs/inputDay4/actualinput4_1.txt"
	inputs2 := []int{4, 9}
	outputs := []string{"CMZ", "GRTSWNJHH"}

	for i := 0; i < len(inputs); i++ {
		out := Puzzle5_1(inputs[i], inputs2[i])
		if out != outputs[i] {
			t.Errorf("Expected %v, got %v", outputs[i], out)
		}
	}

}

func TestPuzzle5_2(t *testing.T) {
	inputs := []string{"inputs/inputDay5/input5_1.txt", "inputs/inputDay5/actualinput5_1.txt"} //"inputs/inputDay4/actualinput4_1.txt"
	inputs2 := []int{4, 9}
	outputs := []string{"MCD", "QLFQDBBHM"}

	for i := 0; i < len(inputs); i++ {
		out := Puzzle5_2(inputs[i], inputs2[i])
		if out != outputs[i] {
			t.Errorf("Expected %v, got %v", outputs[i], out)
		}
	}

}
