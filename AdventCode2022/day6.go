package main

func Puzzle6_1(filename string) []int {
	lines := ReadFiletoStrArray(filename)
	out := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		out[i] = getUniqueChars(lines[i], 4)
	}
	return out
}

func Puzzle6_2(filename string) []int {
	lines := ReadFiletoStrArray(filename)
	out := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		out[i] = getUniqueChars(lines[i], 14)
	}
	return out
}

func getUniqueChars(line string, cnt int) int {
	max := len(line)
	for i := 0; i < max-cnt; i++ {
		slc := line[i : i+cnt]
		if isuniqueSlice(slc) {
			return i + cnt
		}
	}
	return -1
}

func isuniqueSlice(slc string) bool {
	for i := 0; i < len(slc); i++ {
		for j := 0; j < len(slc); j++ {
			if i != j {
				if slc[i] == slc[j] {
					return false
				}
			}
		}
	}
	return true
}
