package doublylinkedlist

func (dl *DoublyLinkedList) removeNodeAtBeginning() {
	if dl.Length == 0 {
		return
	}

	temp := dl.Head
	if dl.Head == dl.Tail {
		dl.Tail = nil
	} else {
		dl.Head.Next.Previous = nil
	}

	dl.Head = dl.Head.Next
	temp.Next = nil
	dl.Length--
}

func (dl *DoublyLinkedList) removeNodeAtEnd() {
	if dl.Length == 0 {
		return
	}

	if dl.Head == dl.Tail {
		dl.Head = nil
		dl.Tail = nil
		return
	}

	dl.Tail.Previous.Next = nil
	dl.Tail.Previous, dl.Tail = nil, dl.Tail.Previous
	dl.Length--
}
