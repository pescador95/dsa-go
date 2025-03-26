package linked_list

type Node struct {
	Data string
	Next *Node
}

func main() {

	nodes := Node{
		Data: "pescador95",
		Next: &Node{
			Data: "channel",
			Next: &Node{
				Data: "twitch",
				Next: &Node{
					Data: "twitch.tv/pescador95",
					Next: nil,
				},
			},
		},
	}

	printNodes(&nodes)

	nodes2 := &LinkedList{}
	nodes2.insertAtBeginning("pescador95")
	nodes2.insertAtBeginning("channel")
	nodes2.insertAtBeginning("twitch")
	nodes2.insertAtBeginning("twitch.tv/pescador95")
	printNodes(nodes2.Head)

	nodes2.insertAtGivenPosition("youtube", 2)
	printNodes(nodes2.Head)

}
