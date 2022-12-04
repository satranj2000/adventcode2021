package main

import (
	"strconv"
	"strings"
)

func Puzzle4_1(filename string) int {
	cnt := 0
	assgns := ReadFiletoStrArray(filename)
	for i := 0; i < len(assgns); i++ {
		batches := strings.Split(assgns[i], ",")
		batch1 := strings.Split(batches[0], "-")
		b1_1, _ := strconv.Atoi(batch1[0])
		b1_2, _ := strconv.Atoi(batch1[1])

		batch2 := strings.Split(batches[1], "-")
		b2_1, _ := strconv.Atoi(batch2[0])
		b2_2, _ := strconv.Atoi(batch2[1])

		//2-8,3-7; 2-6,4-0
		if b2_1 >= b1_1 && b2_2 <= b1_2 {
			cnt++
		} else {
			//3-7, 2-8 ; 6-6, 4-6
			if b1_1 >= b2_1 && b1_2 <= b2_2 {
				cnt++
			}

		}
	}
	return cnt
}

func Puzzle4_2(filename string) int {
	cnt := 0
	assgns := ReadFiletoStrArray(filename)
	tot := 0
	for i := 0; i < len(assgns); i++ {
		batches := strings.Split(assgns[i], ",")
		batch1 := strings.Split(batches[0], "-")
		b1_1, _ := strconv.Atoi(batch1[0])
		b1_2, _ := strconv.Atoi(batch1[1])

		batch2 := strings.Split(batches[1], "-")
		b2_1, _ := strconv.Atoi(batch2[0])
		b2_2, _ := strconv.Atoi(batch2[1])

		//fmt.Printf("%d,%d - %d,%d", b1_1, b1_2, b2_1, b2_2)
		tot++
		//2-5,4-6; 6-6,4-6; 2-8,3-7; 1-1,1-4; 4-4,4-4; 7-97; 3-6
		if b1_2 < b2_1 {
			cnt++
		} else {
			if b2_2 < b1_1 {
				cnt++
			}
		}
	}
	return tot - cnt
}
