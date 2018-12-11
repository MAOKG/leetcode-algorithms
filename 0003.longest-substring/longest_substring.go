package problem3

func lengthOfLongestSubstring(s string) int {
	if len(s) == 1 {
		return 1
	}
	if len(s) == 0 {
		return 0
	}
	var i, j int
	max := 1
	r := []rune(s)
	// map the last seen location
	locationMap := make(map[rune]int)
	// slide window [i, j)
	for i < len(r) && j < len(r) {
		_, ok := locationMap[r[j]]
		if ok {
			delete(locationMap, r[i])
			i++
		} else {
			locationMap[r[j]] = j
			j++
			if j-i > max {
				max = j - i
			}
		}
	}
	return max
}
