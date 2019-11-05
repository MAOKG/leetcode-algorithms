package problem215

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	// case 1
	input := []int{3, 2, 1, 5, 6, 4}
	k := 2
	result := findKthLargest(input, k)
	assert.Equal(t, 5, result)

	// case 2
	input = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	result = findKthLargest(input, k)
	assert.Equal(t, 4, result)
}
