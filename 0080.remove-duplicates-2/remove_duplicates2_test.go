package problem80

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	// case 1
	nums := []int{1, 1, 1, 2, 2, 3}
	l := removeDuplicates(nums)
	assert.Equal(t, 5, l)
	assert.Equal(t, []int{1, 1, 2, 2, 3}, nums[:5])

	// case 2
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	l = removeDuplicates(nums)
	assert.Equal(t, 7, l)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 3, 3}, nums[:7])

	// case 3
	nums = []int{0, 0}
	l = removeDuplicates(nums)
	assert.Equal(t, 2, l)
	assert.Equal(t, []int{0, 0}, nums[:2])

	// case 4
	nums = []int{1, 2, 2}
	l = removeDuplicates(nums)
	assert.Equal(t, 3, l)
	assert.Equal(t, []int{1, 2, 2}, nums[:3])

	// case 5
	nums = []int{1, 1, 2, 2}
	l = removeDuplicates(nums)
	assert.Equal(t, 4, l)
	assert.Equal(t, []int{1, 1, 2, 2}, nums[:4])
}
