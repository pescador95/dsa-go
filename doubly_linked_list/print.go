package doublylinkedlist

import "fmt"

func (dl *DoublyLinkedList) printList() {
	current := dl.Head

	for current != nil {
		fmt.Printf("%d ->", current.Data)
		current = current.Next
	}

	fmt.Print("null")
}

func (ll *DoublyLinkedList) printListBackWard() {

	current := ll.Tail

	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Previous
	}

	fmt.Printf("null\n")
}
