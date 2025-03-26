package linked_list

func (ll *LinkedList) reverseList() {
	current := ll.Head
	var prev *Node
	var next *Node

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	ll.Head = prev
}
