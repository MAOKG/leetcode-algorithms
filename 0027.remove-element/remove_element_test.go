package problem27

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	// case 1
	nums1 := []int{3, 2, 2, 3}
	val := 3

	l := removeElement(nums1, val)
	assert.Equal(t, 2, l)
	assert.ElementsMatch(t, []int{2, 2}, nums1[:2])

	// case 2
	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2

	l = removeElement(nums2, val)
	assert.Equal(t, 5, l)
	assert.ElementsMatch(t, []int{0, 1, 3, 0, 4}, nums2[:5])

	// case 3
	nums3 := []int{0}
	val = 0

	l = removeElement(nums3, val)
	assert.Equal(t, 0, l)

	// case 4
	nums4 := []int{2, 2}
	val = 2

	l = removeElement(nums4, val)
	assert.Equal(t, 0, l)
}
