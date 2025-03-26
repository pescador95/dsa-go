package linked_list

import "fmt"

func printNodes(n *Node) {
	for n != nil {
		print(n.Data, " --> ")
		n = n.Next
	}
	fmt.Println("null")
}
