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
		return float64(nums1[0]+nums2[2]) / 2
	}
	// isEven := (l1+l2)%2 == 0
	if l1 == 1 {
		if nums1[0] <= nums2[l1/2-1] {
			return float64(nums2[l1/2-1])
		}
		if nums1[0] <= nums2[l1/2] {
			return float64(nums1[0])
		}
		return float64(nums2[l1/2])
	}
	if l2 == 1 {
		if nums2[0] <= nums1[l1/2-1] {
			return float64(nums1[l1/2-1])
		}
		if nums2[0] <= nums1[l1/2] {
			return float64(nums2[0])
		}
		return float64(nums1[l1/2])
	}
	if nums1[l1-1] <= nums2[0] {
		if l1 < l2 {
			return float64(nums2[(l2+l1)/2-l1])
		}
		return float64(nums1[(l1+l2)/2])
	}
	if nums2[l2-1] <= nums1[0] {
		if l2 < l1 {
			return float64(nums1[(l2+l1)/2-l2])
		}
		return float64(nums2[(l1+l2)/2])
	}

	k1 := findMedian(nums1, nums2, l1, l2, 0, l1-1)
	if k1 > -1 {
		return float64(nums1[k1])
	}
	// if k1 >= 0 {
	// 	m1 := nums1[k1]
	// 	if isEven {
	// 		var m2 int
	// 		if k1 < l1-1 {
	// 			a := nums1[k1+1]
	// 			b := nums2[(l1+l2)/2-k1]
	// 			if a <= b {
	// 				m2 = a
	// 			} else {
	// 				m2 = b
	// 			}
	// 		} else {
	// 			b := nums2[(l1+l2)/2-k1]
	// 			m2 = b
	// 		}
	// 		return float64(m1+m2) / 2
	// 	}
	// 	return float64(m1)
	// }
	k2 := findMedian(nums2, nums1, l2, l1, 0, l2-1)

	// m1 := nums2[k2]
	// if isEven {
	// 	var m2 int
	// 	if k1 < l1-1 {
	// 		a := nums2[k2+1]
	// 		b := nums1[(l1+l2)/2-k2]
	// 		if a <= b {
	// 			m2 = a
	// 		} else {
	// 			m2 = b
	// 		}
	// 	} else {
	// 		b := nums1[(l1+l2)/2-k2]
	// 		m2 = b
	// 	}
	// 	return float64(m1+m2) / 2
	// }
	return float64(nums2[k2])
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

func findMedian(nums1 []int, nums2 []int, l1 int, l2 int, low int, heigh int) int {
	if low > heigh {
		return -1
	}
	k := (low + heigh) / 2
	if k <= (l1-l2)/2 {
		return findMedian(nums1, nums2, l1, l2, k+1, heigh)
	}
	if k >= (l1+l2)/2 {
		return findMedian(nums1, nums2, l1, l2, low, k-1)
	}
	if nums1[k] >= nums2[(l1+l2)/2-k-1] && nums1[k] <= nums2[(l1+l2)/2-k] {
		return k
	}
	if nums1[k] > nums2[(l1+l2)/2-k] {
		return findMedian(nums1, nums2, l1, l2, low, k-1)
	}
	return findMedian(nums1, nums2, l1, l2, k+1, heigh)
}

/* thought process
assume nums1[k] is the median
=>
in nums1, k elements <= nums1[k], l1-k elements >= nums1[k]
in nums2, (l1+l2)/2 - k elements <= nums1[k], (l2-l1)/2 +k elements >= nums1[k]
=>
exist k such that nums2[(l1+l2)/2 - k - 1] <= nums1[k] <= nums2[(l1+l2)/2 - k]
 (l1+l2)/2 -k < l2 => (l1- l2) < 2k
 (l1+l2)/2 -k >0 => (l1+l2) > 2k
**/
