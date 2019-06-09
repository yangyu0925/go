package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

func CreateTree(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int)  {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	node.Value = value
}

