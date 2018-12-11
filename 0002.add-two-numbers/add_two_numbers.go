package problem2

// ListNode for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	currNode := dummyHead
	carry := 0
	a, b := l1, l2

	for a != nil || b != nil || carry > 0 {
		val := carry
		if a != nil {
			val += a.Val
			a = a.Next
		}
		if b != nil {
			val += b.Val
			b = b.Next
		}
		currNode.Next = &ListNode{Val: val % 10}
		currNode = currNode.Next
		carry = val / 10
	}

	return dummyHead.Next
}

// 28ms, 24%
