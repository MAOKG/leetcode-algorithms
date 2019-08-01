package problem80

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	l := 1
	curr := nums[0]
	isFull := false //if the curr value already duplicated twice

	for i, val := range nums {
		if val != curr {
			nums[l] = val
			curr = val
			isFull = false
			l++
		} else if !isFull && i > 0 {
			nums[l] = val
			isFull = true
			l++
		}
	}

	return l
}

// 4ms, runtime 93.97%, memory 35.71%
