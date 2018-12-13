package problem4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, 2.0, result)

	nums1 = []int{1, 4}
	nums2 = []int{2, 3, 5}
	result = findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, 3.0, result)
}
