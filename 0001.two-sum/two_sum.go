package twosum

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i, num := range nums {
		j, ok := numsMap[target-num]
		if ok {
			return []int{j, i}
		}
		numsMap[num] = i
	}

	return nil
}
