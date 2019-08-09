package problem88

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)
	curr, i, j := 0, 0, 0
	for curr < m+n {
		if i >= m {
			nums3[curr] = nums2[j]
			curr++
			j++
		} else if j >= n {
			nums3[curr] = nums1[i]
			curr++
			i++
		} else if nums1[i] <= nums2[j] {
			nums3[curr] = nums1[i]
			curr++
			i++
		} else {
			nums3[curr] = nums2[j]
			curr++
			j++
		}
	}
	copy(nums1, nums3)
}

// runtime: 0ms, 100%
// memory 36mb, 66.67%
