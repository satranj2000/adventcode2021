package main

import (
	"fmt"
)

func ReadinputdataDay5(filename string, lcnt int) ([]string, []string) {
	strs := ReadFiletoStrArray(filename)
	return strs[:lcnt-1], strs[lcnt+1:]
}

func Puzzle5_1(filename string, brk int) string {

	stacksStrings, instructionStrs := ReadinputdataDay5(filename, brk)
	grid := getArrayListfromStack(stacksStrings)
	//fmt.Printf("%v", grid)
	stacks := getArrayOfStacks(grid)
	//log.Println(stacks)
	instructs := getMovements(instructionStrs)
	//log.Println(instructs)

	output := MoveStacks(stacks, instructs)
	return output
}

func Puzzle5_2(filename string, brk int) string {

	stacksStrings, instructionStrs := ReadinputdataDay5(filename, brk)
	grid := getArrayListfromStack(stacksStrings)
	//fmt.Printf("%v", grid)
	stacks := getArrayOfStacks(grid)
	//log.Println(stacks)
	instructs := getMovements(instructionStrs)
	//log.Println(instructs)

	output := MoveStacks2(stacks, instructs)
	return output
}

func MoveStacks2(stacks []Stack, instructs [][]int) string {

	for i := 0; i < len(instructs); i++ {
		times := instructs[i][0]
		fromStk := instructs[i][1] - 1
		toStk := instructs[i][2] - 1
		orderOfOut := make([]string, times)
		for j := 0; j < times; j++ {
			out, _ := stacks[fromStk].Pop()
			orderOfOut[j] = out
		}
		for k := times; k > 0; k-- {
			stacks[toStk].Push(orderOfOut[k-1])
		}

		fmt.Printf("Stack %v value is %v, after %d\n", fromStk, stacks[fromStk], i)
	}
	outstr := ""
	for i := 0; i < len(stacks); i++ {
		out, _ := stacks[i].Pop()
		outstr += out
		stacks[i].Push(out)
	}

	return outstr
}

func MoveStacks(stacks []Stack, instructs [][]int) string {

	for i := 0; i < len(instructs); i++ {
		times := instructs[i][0]
		fromStk := instructs[i][1] - 1
		toStk := instructs[i][2] - 1
		for j := 0; j < times; j++ {
			out, _ := stacks[fromStk].Pop()
			stacks[toStk].Push(out)
		}
		fmt.Printf("Stack %v value is %v, after %d\n", fromStk, stacks[fromStk], i)
	}
	outstr := ""
	for i := 0; i < len(stacks); i++ {
		out, _ := stacks[i].Pop()
		outstr += out
		stacks[i].Push(out)
	}

	return outstr
}

func getMovements(instructions []string) [][]int {
	instructs := make([][]int, 0, len(instructions))

	for i := 0; i < len(instructions); i++ {
		moves := make([]int, 3)
		fmt.Sscanf(instructions[i], "move %d from %d to %d", &moves[0], &moves[1], &moves[2])
		instructs = append(instructs, moves)
	}
	return instructs
}

func getArrayListfromStack(stacks []string) [][]byte {
	l := len(stacks)
	grid := make([][]byte, len(stacks))
	noofcols := (len(stacks[0]) + 1) / 4
	for i := 0; i < l; i++ {
		grid[i] = make([]byte, noofcols)
	}

	for i := 0; i < len(stacks); i++ {
		line := []byte(stacks[i])
		col := 0
		for j := 0; j < len(line); j += 4 {
			if line[j+1] != ' ' {
				grid[i][col] = line[j+1]
			} else {
				grid[i][col] = ' '
			}
			col++
		}
	}
	return grid
}

func getArrayOfStacks(grids [][]byte) []Stack {
	stacks := []Stack{}
	l := len(grids[0])
	for i := 0; i < l; i++ {
		Stk := Stack{}
		j := len(grids) - 1
		for ; j >= 0; j-- {
			if grids[j][i] != ' ' {
				Stk.Push(string(grids[j][i]))
			} else {
				stacks = append(stacks, Stk)
				break
			}
		}
		if j < 0 {
			stacks = append(stacks, Stk)
		}
	}
	return stacks
}
