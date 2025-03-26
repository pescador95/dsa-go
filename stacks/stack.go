package stacks

type StackLinkedList struct {
	Top    *StackNode
	Length int
}

type StackNode struct {
	Data string
	Next *StackNode
}

func (sl *StackLinkedList) push(item string) {
	newNode := &StackNode{
		Data: item,
		Next: sl.Top,
	}
	sl.Top = newNode
	sl.Length++
}

func (sl *StackLinkedList) pop() string {
	result := sl.Top.Data
	sl.Top = sl.Top.Next
	sl.Length--

	return result
}
