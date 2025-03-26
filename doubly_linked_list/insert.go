package doublylinkedlist

func (dl *DoublyLinkedList) insertNodeAtBeginning(data string) {
	newNode := &ListNode{Data: data}

	if dl.Length == 0 {
		dl.Tail = newNode
	} else {
		dl.Head.Previous = newNode
	}

	newNode.Next = dl.Head
	dl.Head = newNode
	dl.Length++
}

func (dl *DoublyLinkedList) insertNodeAtEnd(data string) {
	newNode := &ListNode{Data: data}

	if dl.Length == 0 {
		dl.Head = newNode
	} else {
		dl.Tail.Next = newNode
		newNode.Previous = dl.Tail
	}

	dl.Tail = newNode
	dl.Length++
}
