package circularlinkedlist

import "fmt"

func (cl *Circularlinkedlist) printList() {

	if cl.Length == 0 {
		return
	}

	first := cl.Last.Next
	for first != cl.Last {
		fmt.Printf("%s -> ", first.Data)
		first = first.Next
	}

	fmt.Printf("%d -> null", first.Data)
}
