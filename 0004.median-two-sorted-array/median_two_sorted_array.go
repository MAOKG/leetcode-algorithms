package problem4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 == 0 && l2 == 0 {
		return 0
	}
	if l1 == 0 {
		return findMedianSingleArray(nums2)
	}
	if l2 == 0 {
		return findMedianSingleArray(nums1)
	}
	if l1 == 1 && l2 == 1 {
		return float64(nums1[0]+nums2[0]) / 2
	}
	if l1 <= l2 {
		return findMedian(nums1, nums2, l1, l2, 0, l1)
	}
	return findMedian(nums2, nums1, l2, l1, 0, l2)
}

// find the median assuming l1 <= l2
func findMedian(nums1 []int, nums2 []int, l1 int, l2 int, low int, heigh int) float64 {
	var maxFirstHalf, minSecondHalf int
	isEven := (l1+l2)%2 == 0
	if l1 == 1 {
		if nums1[0] <= nums2[l2/2] {
			maxFirstHalf = findMax(nums1[0], nums2[l2/2-1])
			if !isEven {
				return float64(maxFirstHalf)
			}
			return float64(maxFirstHalf+nums2[l2/2]) / 2
		}
		maxFirstHalf = nums2[l2/2]
		if !isEven {
			return float64(maxFirstHalf)
		}
		minSecondHalf = findMin(nums1[0], nums2[l2/2+1])
		return float64(maxFirstHalf+minSecondHalf) / 2
	}

	if nums1[l1-1] <= nums2[0] {
		if l1 == l2 {
			return float64(nums1[l1-1]+nums2[0]) / 2
		}
		maxFirstHalf = nums2[(l2-l1+1)/2-1]
		if !isEven {
			return float64(maxFirstHalf)
		}
		minSecondHalf = nums2[(l2-l1+1)/2]
		return float64(maxFirstHalf+minSecondHalf) / 2
	}
	if nums1[0] >= nums2[l2-1] {
		if l1 == l2 {
			return float64(nums1[0]+nums2[l2-1]) / 2
		}
		maxFirstHalf = nums2[(l2+l1+1)/2-1]
		if !isEven {
			return float64(maxFirstHalf)
		}
		minSecondHalf = nums2[(l2+l1+1)/2]
		return float64(maxFirstHalf+minSecondHalf) / 2
	}

	// partition nums1 and nums2
	i := (low + heigh) / 2
	j := (l1+l2+1)/2 - i
	if i >= l1 {
		maxFirstHalf = findMax(nums1[l1-1], nums2[(l2-l1+1)/2-1])
		if !isEven {
			return float64(maxFirstHalf)
		}
		return float64(maxFirstHalf+nums2[(l2-l1+1)/2]) / 2
	}
	if i == 0 {
		maxFirstHalf = nums2[(l1+l2+1)/2-1]
		if !isEven {
			return float64(maxFirstHalf)
		}
		minSecondHalf = findMin(nums1[0], nums2[(l1+l2+1)/2])
		return float64(maxFirstHalf+minSecondHalf) / 2
	}
	if nums1[i-1] > nums2[j] {
		return findMedian(nums1, nums2, l1, l2, low, i)
	}
	if nums2[j-1] > nums1[i] {
		return findMedian(nums1, nums2, l1, l2, i+1, heigh)
	}
	maxFirstHalf = findMax(nums1[i-1], nums2[j-1])
	if !isEven {
		return float64(maxFirstHalf)
	}
	minSecondHalf = findMin(nums1[i], nums2[j])
	return float64(maxFirstHalf+minSecondHalf) / 2
}

// helper function to find median of a single sorted array
func findMedianSingleArray(nums []int) float64 {
	l := len(nums)
	if l%2 == 1 {
		return float64(nums[l/2])
	}
	m := float64(nums[l/2-1]+nums[l/2]) / 2
	return m
}

// helper function to find the min of two integers
func findMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// helper function to find the max of two integers
func findMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// find the partition i, j where nums1[i-1] <= nums2[j] && nums2[j-1] <= num1[i]
