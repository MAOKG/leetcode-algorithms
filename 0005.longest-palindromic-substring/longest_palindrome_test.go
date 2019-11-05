package problem5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	input := "babad"
	output := longestPalindrome(input)
	isExpected := output == "bab" || output == "aba"
	assert.Equal(t, true, isExpected)

	input = "cbbd"
	output = longestPalindrome(input)
	assert.Equal(t, "bb", output)
}
