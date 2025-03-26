package linked_list

func (ll *LinkedList) insertAtBeginning(data string) {
	newNode := &Node{
		Data: data, Next: ll.Head,
	}
	ll.Head = newNode
}

func (ll *LinkedList) insertAtEnd(data string) {
	newNode := &Node{
		Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	last := ll.Head
	for last.Next != nil {
		last = last.Next
	}

	last.Next = newNode
}

func (ll *LinkedList) insertAtGivenPosition(data string, position int) {
	newNode := &Node{
		Data: data,
	}

	if position == 1 {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	previous := ll.Head
	count := 1

	for count < position-1 {
		previous = previous.Next
		count++
	}

	current := previous.Next
	newNode.Next = current
	previous.Next = newNode
}
