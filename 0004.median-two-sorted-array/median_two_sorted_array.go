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

// Helper function to find median of a single sorted array
func findMedianSingleArray(nums []int) float64 {
	l := len(nums)
	if l%2 == 0 {
		return float64(nums[l/2])
	}
	m := float64(nums[l/2]+nums[l/2+1]) / 2
	return m
}

// find the partition i, j where nums1[i-1] <= nums2[j] && nums2[j-1] <= num1[i]
// find the median assuming l1 <= l2
func findMedian(nums1 []int, nums2 []int, l1 int, l2 int, low int, heigh int) float64 {
	var maxFirstHalf, minSecondHalf float64
	isEven := (l1+l2)%2 == 0
	if l1 == 1 {
		if nums1[0] <= nums2[l2/2] {
			maxFirstHalf = float64(nums1[0])
			if nums2[l2/2-1] > nums1[0] {
				maxFirstHalf = float64(nums2[l2/2-1])
			}
			if !isEven {
				return maxFirstHalf
			}
			return (maxFirstHalf + float64(nums2[l2/2])) / 2
		}
		maxFirstHalf = float64(nums2[l2/2])
		if !isEven {
			return maxFirstHalf
		}
		minSecondHalf = float64(nums1[0])
		if nums1[0] >= nums2[l2/2+1] {
			minSecondHalf = float64(nums2[l2/2+1])
		}
		return (maxFirstHalf + minSecondHalf) / 2
	}

	if nums1[l1-1] <= nums2[0] {
		if l1 == l2 {
			return float64(nums1[l1-1]+nums2[0]) / 2
		}
		maxFirstHalf = float64(nums2[(l2-l1+1)/2-1])
		if !isEven {
			return maxFirstHalf
		}
		minSecondHalf = float64(nums2[(l2-l1+1)/2])
		return (maxFirstHalf + minSecondHalf) / 2
	}

	if nums1[0] >= nums2[l2-1] {
		if l1 == l2 {
			return float64(nums1[0]+nums2[l2-1]) / 2
		}
		maxFirstHalf = float64(nums2[(l2+l1+1)/2-1])
		if !isEven {
			return maxFirstHalf
		}
		minSecondHalf = float64(nums2[(l2+l1+1)/2])
		return (maxFirstHalf + minSecondHalf) / 2
	}

	// partition nums1 and nums2
	i := (low + heigh) / 2
	j := (l1+l2+1)/2 - i
	if nums1[i-1] > nums2[j] {
		return findMedian(nums1, nums2, l1, l2, low, i)
	}
	if nums2[j-1] > nums1[i] {
		return findMedian(nums1, nums2, l1, l2, i, heigh)
	}
	maxFirstHalf = float64(nums1[i-1])
	if nums1[i-1] < nums2[j-1] {
		maxFirstHalf = float64(nums2[j-1])
	}
	if !isEven {
		return maxFirstHalf
	}
	minSecondHalf = float64(nums1[i])
	if nums1[i] > nums2[j] {
		minSecondHalf = float64(nums2[j])
	}

	return (maxFirstHalf + minSecondHalf) / 2
}

// find the partition i, j where nums1[i-1] <= nums2[j] && nums2[j-1] <= num1[i]
