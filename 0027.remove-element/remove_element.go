package problem27

func removeElement(nums []int, val int) int {
	l, r, count := 0, len(nums)-1, 0
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
