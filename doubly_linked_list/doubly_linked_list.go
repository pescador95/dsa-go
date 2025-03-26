package doublylinkedlist

type DoublyLinkedList struct {
	Head   *ListNode
	Tail   *ListNode
	Length int
}

type ListNode struct {
	Data     string
	Previous *ListNode
	Next     *ListNode
}
