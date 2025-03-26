package leetcode

func hasCycle(head *ListNode) bool {
	fastPointer := head
	slowPointer := head

	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next

		if fastPointer == slowPointer {
			return true
		}
	}

	return false
}
