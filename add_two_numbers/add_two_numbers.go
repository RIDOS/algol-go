package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers returns the sum of two numbers represented by linked lists l1 and l2.
// The linked lists are non-empty and each Linked List represents a non-negative integer.
// The ones place of the two numbers is at the head of each linked list, and each node
// contains a single digit. The function adds the two numbers and returns the sum as a
// linked list.
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	newNode := &ListNode{}
	curr := newNode
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		val := sum % 10

		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return newNode.Next
}
