package problem2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	// 2 -> 4 -> 3
	a1 := ListNode{3, nil}
	a2 := ListNode{4, &a1}
	a3 := ListNode{2, &a2}

	// 5 -> 6 -> 4
	b1 := ListNode{4, nil}
	b2 := ListNode{6, &b1}
	b3 := ListNode{5, &b2}

	// 7 -> 0 -> 8
	c1 := ListNode{8, nil}
	c2 := ListNode{0, &c1}
	c3 := ListNode{7, &c2}

	result := addTwoNumbers(&a3, &b3)
	assert.Equal(t, &c3, result)

	// 5 + 5 = 0 -> 1
	a := ListNode{5, nil}
	b := ListNode{5, nil}
	c1 = ListNode{1, nil}
	c2 = ListNode{0, &c1}
	result = addTwoNumbers(&a, &b)
	assert.Equal(t, &c2, result)

	// 0 + 8 -> 1 = 8 -> 1
	a = ListNode{0, nil}
	b1 = ListNode{1, nil}
	b2 = ListNode{8, &b1}
	c1 = ListNode{1, nil}
	c2 = ListNode{8, &c1}
	result = addTwoNumbers(&a, &b2)
	assert.Equal(t, &c2, result)
}
