package circularlinkedlist

func (cl *Circularlinkedlist) insertNodeAtBeginning(data string) {
	newNode := &ListNode{
		Data: data,
	}

	if cl.Length == 0 || cl.Last == nil {
		cl.Last = newNode
	} else {
		newNode.Next = cl.Last.Next
	}

	cl.Last.Next = newNode
	cl.Length++
}

func (cl *Circularlinkedlist) insertNodeAtEnd(data string) {
	newNode := &ListNode{
		Data: data,
	}

	if cl.Length == 0 || cl.Last == nil {
		cl.Last = newNode
		cl.Last.Next = cl.Last
	} else {
		newNode.Next, cl.Last.Next = cl.Last.Next, newNode
		cl.Last = newNode
	}

	cl.Length++
}
