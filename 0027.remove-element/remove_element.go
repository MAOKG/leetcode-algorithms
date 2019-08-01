package problem27

func removeElement(nums []int, val int) int {
	count := 0
	l := 0
	r := len(nums) - 1
	for l <= r {
		if nums[l] == val {
			if nums[r] != val {
				nums[l] = nums[r]
				l++
			}
			r--
			count++
		} else {
			l++
		}
	}

	return len(nums) - count
}

// 0ms, 61.54%
