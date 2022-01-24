package linkedList

func hasCycle(head *ListNode) bool {
	var (
		fast = head
		slow = head
	)
	if slow == nil || slow.Next == nil {
		return false
	}

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
