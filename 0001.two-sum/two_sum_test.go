package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expectedResult := []int{0, 1}
	result := twoSum(nums, target)
	assert.Equal(t, expectedResult, result)
}
