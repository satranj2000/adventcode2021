package main

import (
	"strconv"
)

func Puzzle8_1(filename string) int {

	lines := ReadFiletoStrArray(filename)
	maptrees := readfiletoMatrix(lines)

	cnt := getCountofvisibleTrees(maptrees)

	return cnt
}

func Puzzle8_2(filename string) int {

	lines := ReadFiletoStrArray(filename)
	maptrees := readfiletoMatrix(lines)

	spot := getScenicSpot(maptrees)

	return spot

}

func getScenicSpot(trees [][]int) int {
	maxr := len(trees) - 1
	maxc := len(trees[0]) - 1
	maxpos := 0
	for i := 1; i < maxr; i++ {
		for j := 1; j < maxc; j++ {
			lr := checkVisibilityLeftRow(trees, i, j, trees[i][j])
			rr := checkVisibilityRightRow(trees, i, j, trees[i][j])
			tc := checkVisibilityTopCol(trees, j, i, trees[i][j])
			bc := checkVisibilityBottomCol(trees, j, i, trees[i][j])

			newval := lr * rr * tc * bc
			if newval > maxpos {
				maxpos = newval
			}
		}
	}

	return maxpos

}

func checkVisibilityLeftRow(trees [][]int, rowid int, pos int, compare int) int {
	cnt := 0
	for j := pos - 1; j >= 0; j-- {
		if trees[rowid][j] < compare {
			cnt++
		} else if trees[rowid][j] >= compare {
			cnt++
			break
		}
	}
	return cnt
}

func checkVisibilityRightRow(trees [][]int, rowid int, pos int, compare int) int {
	cnt := 0
	for j := pos + 1; j < len(trees[rowid]); j++ {
		if trees[rowid][j] < compare {
			cnt++
		} else if trees[rowid][j] >= compare {
			cnt++
			break
		}
	}
	return cnt
}

func checkVisibilityTopCol(trees [][]int, colid int, pos int, compare int) int {
	cnt := 0
	for i := pos - 1; i >= 0; i-- {
		if trees[i][colid] < compare {
			cnt++
		} else if trees[i][colid] >= compare {
			cnt++
			break
		}
	}
	return cnt
}

func checkVisibilityBottomCol(trees [][]int, colid int, pos int, compare int) int {
	cnt := 0
	for i := pos + 1; i < len(trees); i++ {
		if trees[i][colid] < compare {
			cnt++
		} else if trees[i][colid] >= compare {
			cnt++
			break
		}
	}
	return cnt
}

func getCountofvisibleTrees(trees [][]int) int {

	maxr := len(trees) - 1
	maxc := len(trees[0]) - 1
	count := 0
	for i := 0; i <= maxr; i++ {
		for j := 0; j <= maxc; j++ {
			if i == 0 || i == maxr || j == 0 || j == maxc {
				count++
				continue
			}
			isVisibleRow := checkRow(trees, i, j, trees[i][j])
			if isVisibleRow {
				count++
				continue
			} else {
				isVisibleCol := checkCol(trees, j, i, trees[i][j])
				if isVisibleCol {
					count++
					continue
				}
			}
		}

	}
	return count
}

func checkRow(trees [][]int, rowid int, pos int, compare int) bool {
	checkOtherside := false
	for i := 0; i < pos; i++ {
		if trees[rowid][i] >= compare {
			checkOtherside = true
			break
		}
	}
	if checkOtherside {
		rowlen := len(trees[rowid])
		for i := pos + 1; i < rowlen; i++ {
			if trees[rowid][i] >= compare {
				return false
			}
		}
	}

	return true
}

func checkCol(trees [][]int, colid int, pos int, compare int) bool {
	checkOtherside := false
	for i := 0; i < pos; i++ {
		if trees[i][colid] >= compare {
			checkOtherside = true
			break
		}
	}
	if checkOtherside {
		rowscnt := len(trees)
		for i := pos + 1; i < rowscnt; i++ {
			if trees[i][colid] >= compare {
				return false
			}
		}
	}

	return true
}

func readfiletoMatrix(lines []string) [][]int {
	values := make([][]int, 0, len(lines))

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		newline := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			newline[i], _ = strconv.Atoi(string(line[i]))
		}
		values = append(values, newline)
	}
	return values
}
