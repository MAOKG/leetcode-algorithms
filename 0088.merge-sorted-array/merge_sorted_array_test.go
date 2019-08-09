package problem88

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	// case 1
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1[:6])

	// case 2
	nums1 = []int{1, 2, 8, 0, 0, 0}
	m = 3
	nums2 = []int{2, 5, 6}
	n = 3

	merge(nums1, m, nums2, n)
	assert.Equal(t, []int{1, 2, 2, 5, 6, 8}, nums1[:6])
}
