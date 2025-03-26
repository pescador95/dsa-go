package linked_list

func (ll *LinkedList) removeFirstNode() {
	if ll.Head == nil {
		return
	}

	ll.Head = ll.Head.Next
	return
}

func (ll *LinkedList) removeNodeAtEnd() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	previous := ll.Head
	for previous.Next.Next != nil {
		previous = previous.Next
	}

	previous.Next = nil
	return

}

func (ll *LinkedList) removeNodeAtGivenPosition(position int) {
	if ll.Head == nil || position <= 0 {
		return
	}

	if position == 1 {
		ll.Head = ll.Head.Next
		return
	}

	count := 1
	previous := ll.Head

	for count < position-1 {
		previous = previous.Next
		count++
	}

	current := previous.Next
	previous.Next = current.Next
	current = nil
}

func (ll *LinkedList) removeLoop() {
	fastPointer := ll.Head
	slowPointer := ll.Head

	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next

		if slowPointer == fastPointer {
			temp := ll.Head

			for slowPointer.Next != temp.Next {
				temp = temp.Next
				slowPointer = slowPointer.Next
			}

			slowPointer.Next = nil
		}
	}
}
