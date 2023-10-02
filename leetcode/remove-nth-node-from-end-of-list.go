package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	x := len - n
	if x < 0 {
		return nil
	}
	var prev *ListNode
	current = head
	for i := 0; i < len && current != nil; i++ {
		if x == i {
			if prev == nil {
				return head.Next
			}
			prev.Next = current.Next
		}
		if prev == nil {
			prev = head
		} else {
			prev = prev.Next
		}
		current = current.Next
	}
	return head
}
