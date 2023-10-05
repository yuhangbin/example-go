package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := head.Next
	head.Next = result.Next
	result.Next = head
	current := result.Next
	for current.Next != nil && current.Next.Next != nil {
		left := current.Next
		right := current.Next.Next
		// swap
		left.Next = right.Next
		right.Next = left
		current.Next = right
		// update pointer
		current = left
	}
	return result
}
