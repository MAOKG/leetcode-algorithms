package problem215

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return randomizedSelect(nums, 0, n-1, n-k+1)
}

func randomizedSelect(nums []int, l int, r int, i int) int {
	if l == r {
		return nums[l]
	}
	nums1, q := randomizedPartition(nums, l, r)
	k := q - l + 1
	if i == k {
		return nums1[q]
	}
	if i < k {
		return randomizedSelect(nums1, l, q-1, i)
	}
	return randomizedSelect(nums1, q, r, i-k+1)
}

func randomizedPartition(nums []int, l int, r int) ([]int, int) {
	i := random(l, r)
	nums[i], nums[r] = nums[r], nums[i]
	return partiton(nums, l, r)
}

func partiton(nums []int, l int, r int) ([]int, int) {
	x := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return nums, i + 1
}

func random(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	result := min + rand.Intn(max-min+1)
	return result
}

// runtime: 4ms, 99.12%
// memory: 3.6 MB, 50%
