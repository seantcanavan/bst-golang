package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value int
}

var tree1 Node
var tree2 Node
var tree3 Node

func printPreOrder(tree Node) {
	fmt.Println(tree.value)

	if tree.left != nil {
		printPreOrder(*tree.left)
	}

	if tree.right != nil {
		printPreOrder(*tree.right)
	}

}

func printInOrder(tree Node) {
	if tree.left != nil {
		printInOrder(*tree.left)
	}

	fmt.Println(tree.value)

	if tree.right != nil {
		printInOrder(*tree.right)
	}
}

func printPostOrder(tree Node) {
	if tree.left != nil {
		printPostOrder(*tree.left)
	}

	if tree.right != nil {
		printPostOrder(*tree.right)
	}

	fmt.Println(tree.value)
}

func initializeTree1() {
	eight := Node{
		value: 8,
	}

	three := Node{
		value: 3,
	}

	ten := Node{
		value: 10,
	}

	eight.left = &three
	eight.right = &ten

	one := Node{
		value: 1,
	}

	six := Node{
		value: 6,
	}

	three.left = &one
	three.right = &six

	four := Node{
		value: 4,
	}

	seven := Node{
		value: 7,
	}

	six.left = &four
	six.right = &seven

	fourteen := Node{
		value: 14,
	}

	ten.right = &fourteen

	thirteen := Node{
		value: 13,
	}

	fourteen.left = &thirteen

	tree1 = eight
}

func initializeTree2() {
	one := Node{
		value: 1,
	}

	two := Node{
		value: 2,
	}

	four := Node{
		value: 4,
	}

	five := Node{
		value: 5,
	}

	three := Node{
		value: 3,
	}

	one.left = &two
	one.right = &three

	two.left = &four
	two.right = &five

	tree2 = one
}

func main() {
	initializeTree1()
	initializeTree2()

	fmt.Println("preOrder")
	printPreOrder(tree1)

	fmt.Println("postOrder")
	printPostOrder(tree1)

	fmt.Println("inOrder")
	printInOrder(tree1)
}
