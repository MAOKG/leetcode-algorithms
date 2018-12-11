package problem3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubstring(t *testing.T) {
	input := "abcabcbb"
	result := lengthOfLongestSubstring(input)
	assert.Equal(t, 3, result)

	input = "bbbbb"
	result = lengthOfLongestSubstring(input)
	assert.Equal(t, 1, result)

	input = "pwwkew"
	result = lengthOfLongestSubstring(input)
	assert.Equal(t, 3, result)
}
