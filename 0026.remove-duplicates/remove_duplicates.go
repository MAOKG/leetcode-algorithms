package problem26

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	l := 1
	curr := nums[0]

	for _, val := range nums {
		if val != curr {
			nums[l] = val
			curr = val
			l++
		}
	}

	return l
}

// 44ms, runtime 88.5%, memory 23.58%
