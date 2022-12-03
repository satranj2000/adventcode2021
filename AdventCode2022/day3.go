package main

func Puzzle3_1(filename string) int {

	rucksacks := ReadFiletoStrArray(filename)
	cnt := 0
	for i := 0; i < len(rucksacks); i++ {
		val := GetCommonItem(rucksacks[i])
		if val >= 97 {
			cnt += ((int(val) - 97) + 1) // 97 - for Capital letters
		} else {
			cnt += (26 + (int(val) - 65) + 1) // 65 for small letter
		}
	}
	return cnt
}

func Puzzle3_2(filename string) int {
	rucksacks := ReadFiletoStrArray(filename)
	cnt := 0
	for i := 0; i < len(rucksacks); i += 3 {
		cnt += GetCommonAmongThree(rucksacks, i)
	}
	return cnt
}

func GetCommonAmongThree(rucksacks []string, i int) int {
	firStr := GetPosition(rucksacks[i])
	SecStr := GetPosition(rucksacks[i+1])
	ThrStr := GetPosition(rucksacks[i+2])

	for i := 0; i < 52; i++ {
		if firStr[i] && SecStr[i] && ThrStr[i] {
			return i + 1
		}
	}
	return -1
}

func GetPosition(item string) [52]bool {
	posArr := [52]bool{}
	for i := 0; i < len(item); i++ {
		val := item[i]
		pos := 0
		if val >= 97 {
			pos = (int(val) - 97) // 97 - for Capital letters
		} else {
			pos = (26 + (int(val) - 65)) // 65 for small letter
		}
		posArr[pos] = true
	}
	return posArr
}

func GetCommonItem(items string) byte {

	len := len(items) / 2

	//println("%d, %d", 'a', 'A') // 65 - A ; 97 - a
	for i := 0; i < len; i++ {
		for j := len; j < len*2; j++ {
			if items[i] == items[j] {
				println("%d", items[i])
				return items[i]
			}
		}
	}

	return 0
}
