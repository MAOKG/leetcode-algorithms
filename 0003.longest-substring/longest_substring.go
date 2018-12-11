package problem3

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	if len(s) == 0 {
		return 0
	}
	max := 1
	r := []rune(s)
	// map the last seen location
	locationMap := make(map[rune]int)

	for i := 0; i+max < len(s); i++ {
		locationMap[r[i]] = i
		l := 1
		for i+l < len(s) {
			j, ok := locationMap[r[i+l]]
			if ok && j >= i && j < i+l {
				break
			}
			locationMap[r[i+l]] = i + l
			l++
		}
		if l > max {
			max = l
		}
	}

	return max
}
