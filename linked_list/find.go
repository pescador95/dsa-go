package linked_list

func (ll *LinkedList) findElement(data string) int {

	current := ll.Head
	count := 1

	for current != nil {
		if current.Data == data {
			return count
		}
		count++
		current = current.Next
	}

	return -1
}

func (ll *LinkedList) findWhereLoopStarts() *Node {
	fastPointer := ll.Head
	slowPointer := ll.Head

	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		slowPointer = slowPointer.Next

		if slowPointer == fastPointer {
			temp := ll.Head

			for slowPointer != temp {
				temp = temp.Next
				slowPointer = slowPointer.Next
			}

			return temp
		}
	}

	return nil
}

func findNodeLength(n *Node) (count int) {
	for n != nil {
		count++
		n = n.Next
	}
	return
}
