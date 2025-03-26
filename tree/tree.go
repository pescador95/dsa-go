package tree

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func insert(node *TreeNode, data int) *TreeNode {
	if node == nil {
		return &TreeNode{Data: data}
	}
	if data <= node.Data {
		node.Left = insert(node.Left, data)
	} else {
		node.Right = insert(node.Right, data)
	}
	return node
}

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	println(node.Data)
	preOrder(node.Left)
	preOrder(node.Right)
}
