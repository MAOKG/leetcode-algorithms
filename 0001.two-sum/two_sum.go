package problem1

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i, num := range nums {
		j, ok := numsMap[target-num]
		if ok {
			// Need to reverse the order since j appeared before i
			return []int{j, i}
		}
		numsMap[num] = i
	}

	return nil
}
