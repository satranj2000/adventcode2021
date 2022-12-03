package main

import "testing"

func TestPuzzle3_1(t *testing.T) {

	inputs := []string{"inputs/inputDay3/input3_1.txt", "inputs/inputDay3/actualinput3_1.txt"}
	expectedResults := []int{157, 0}
	for i := 0; i < len(inputs); i++ {
		res := Puzzle3_1(inputs[i])
		if res != expectedResults[i] {
			t.Errorf("Expected %d, got %d", expectedResults[i], res)
		}
	}
}

func TestGetCommonItem(t *testing.T) {

	inputs := []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw", "aBbDbxyz"}
	expectedResults := []byte{'p', 'L', 'P', 'v', 't', 's', 'b'}
	for i := 0; i < len(inputs); i++ {
		res := GetCommonItem(inputs[i])
		if res != expectedResults[i] {
			t.Errorf("Expected %d, got %d", expectedResults[i], res)
		}
	}
}

func TestGetPosition(t *testing.T) {

	inputs := []string{"abc"}
	Expoutput := [52]bool{}
	Expoutput[0] = true
	Expoutput[1] = true
	Expoutput[2] = true

	output := GetPosition(inputs[0])

	if output != Expoutput {
		t.Error("Expected ", Expoutput, "Got ", output)
	}

	inputs2 := []string{"azAZ"}
	Expoutput2 := [52]bool{}
	Expoutput2[0] = true
	Expoutput2[25] = true
	Expoutput2[26] = true
	Expoutput2[51] = true
	output2 := GetPosition(inputs2[0])
	if output2 != Expoutput2 {
		t.Error(" For 2nd test - Expected ", Expoutput2, "Got ", output2)
	}
}

func TestGetCommonAmongThree(t *testing.T) {
	inputs := []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg"}

	out := GetCommonAmongThree(inputs, 0)

	if out != 18 {
		t.Errorf("Expected 18, got %d", out)
	}
}

func TestGetCommonAmongThree2(t *testing.T) {
	inputs := []string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw"}

	out := GetCommonAmongThree(inputs, 0)

	if out != 52 {
		t.Errorf("Expected 52, got %d", out)
	}
}

func TestPuzzle3_2(t *testing.T) {

	inputs := []string{"inputs/inputDay3/input3_1.txt", "inputs/inputDay3/actualinput3_1.txt"}
	outputs := []int{70, 0}

	for i := 0; i < len(inputs); i++ {
		out := Puzzle3_2(inputs[i])
		if out != outputs[i] {
			t.Errorf("Got %d, Expected %d", out, outputs[i])
		}
	}
}
