package problem2

// ListNode for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nodes []*ListNode
	var carry int
	a, b := l1, l2

	for (a != nil) || (b != nil) || carry > 0 {
		val := carry
		carry = 0
		if a != nil {
			val += a.Val
			a = a.Next
		}
		if b != nil {
			val += b.Val
			b = b.Next
		}
		if val >= 10 {
			val -= 10
			carry = 1
		}
		nodes = append(nodes, &ListNode{Val: val})
	}

	for index, node := range nodes {
		if index < len(nodes)-1 {
			node.Next = nodes[index+1]
		}
	}

	return nodes[0]
}
