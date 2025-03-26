package circularlinkedlist

type Circularlinkedlist struct {
	Length int
	Last   *ListNode
}

type ListNode struct {
	Data string
	Next *ListNode
}
