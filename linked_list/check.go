package linked_list

func (ll *LinkedList) checkLoop() bool {
	if ll.Head == nil {
		return false
	}

	slow := ll.Head
	fast := ll.Head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}

	return false
}
