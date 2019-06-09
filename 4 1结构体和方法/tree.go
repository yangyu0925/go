package main

import "fmt"

type tree struct {
	value int
	left, right *tree
}

func createTree(value int) *tree {
	return &tree{value: value}
}

func (node tree) print()  {
	fmt.Println(node.value)
}

func (node *tree) setValue (value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.value = value
}

func (node *tree) traverse()  {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root tree

	root = tree{value: 3}
	root.left = &tree{}
	root.right = &tree{5, nil, nil}
	root.right.left = new(tree)
	root.left.right = createTree(2)

	root.print()
	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println()

	//var pRoot *tree
	//pRoot.setValue(200)
	//pRoot = &root
	//pRoot.setValue(300)
	//pRoot.print()

	root.traverse()
}
