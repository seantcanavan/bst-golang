package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value int
}

var tree1 *Node
var tree2 *Node
var tree3 *Node

func printPreOrder(tree *Node) {
	if tree == nil {
		return
	}

	fmt.Println(tree.value)

	if tree.left != nil {
		printPreOrder(tree.left)
	}

	if tree.right != nil {
		printPreOrder(tree.right)
	}

}

func printInOrder(tree *Node) {
	if tree == nil {
		return
	}

	if tree.left != nil {
		printInOrder(tree.left)
	}

	fmt.Println(tree.value)

	if tree.right != nil {
		printInOrder(tree.right)
	}
}

func printPostOrder(tree *Node) {
	if tree == nil {
		return
	}

	if tree.left != nil {
		printPostOrder(tree.left)
	}

	if tree.right != nil {
		printPostOrder(tree.right)
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

	tree1 = &eight
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

	three.left = &two
	three.right = &four

	two.left = &one
	four.right = &five

	tree2 = &three
}

func initializeTree3() {
	fifty := Node{
		value: 50,
	}

	twentyFive := Node{
		value: 25,
	}

	seventyFive := Node{
		value: 75,
	}

	twelve := Node{
		value: 12,
	}

	thirtySix := Node{
		value: 36,
	}

	sixtyThree := Node{
		value: 63,
	}

	eightySeven := Node{
		value: 87,
	}

	six := Node{
		value: 6,
	}

	eighteen := Node{
		value: 18,
	}

	thirty := Node{
		value: 30,
	}

	fortyTwo := Node{
		value: 42,
	}

	fiftySeven := Node{
		value: 57,
	}

	sixtyNine := Node{
		value: 69,
	}

	eightyOne := Node{
		value: 81,
	}

	ninetyThree := Node{
		value: 93,
	}

	fifty.left = &twentyFive
	fifty.right = &seventyFive

	twentyFive.left = &twelve
	twentyFive.right = &thirtySix

	twelve.left = &six
	twelve.right = &eighteen

	thirtySix.left = &thirty
	thirtySix.right = &fortyTwo

	seventyFive.left = &sixtyThree
	seventyFive.right = &eightySeven

	sixtyThree.left = &fiftySeven
	sixtyThree.right = &sixtyNine

	eightySeven.left = &eightyOne
	eightySeven.right = &ninetyThree

	tree3 = &fifty
}

func find(value int, node *Node) *Node {
	// base case and case for value missing
	if node == nil || node.value == value {
		return node
	}

	if node.value > value {
		return find(value, node.left)
	}

	return find(value, node.right)
}

func insert(value int, node *Node) *Node {
	if node.value > value {
		if node.left == nil {
			node.left = &Node{
				value: value,
			}
		}

		return insert(value, node.left)
	}

	if node.value < value {
		if node.right == nil {
			node.right = &Node{
				value: value,
			}
		}

		return insert(value, node.right)
	}

	// value already exists OR tree is empty
	node.value = value
	return node
}

func delete(value int, node *Node) *Node {
	if node == nil { // value wasn't found
		return node
	} else if node.value > value { // go left
		node.left = delete(value, node.left)
	} else if node.value < value { // go right
		node.right = delete(value, node.right)
	} else {
		if node.left == nil && node.right == nil { // if leaf
			return nil
		} else if node.right == nil {
			return node.left
		} else if node.left == nil {
			return node.right
		} else {
			min := minOfTheMax(node.right)
			node.value = min.value
			node.right = delete(min.value, node.right)
		}
	}

	return node
}

func minOfTheMax(right *Node) *Node {
	min := right

	for min.left != nil {
		min = min.left
	}

	return min
}

func maxOfTheMin(left *Node) *Node {
	max := left

	for max.right != nil {
		max = max.right
	}

	return max
}

func main() {
	initializeTree1()
	initializeTree2()
	initializeTree3()

	//fmt.Println("preOrder")
	//printPreOrder(tree3)
	//
	//fmt.Println("postOrder")
	//printPostOrder(tree3)
	//
	//fmt.Println("inOrder")
	//printInOrder(tree3)

	//result := find(100, &tree3)
	//fmt.Println(result)
	//
	//insert(100, &tree3)
	//printInOrder(tree3)
	//
	//insert(125, &tree3)
	//printInOrder(tree3)
	//
	//insert(20, &tree3)
	//printInOrder(tree3)

	tree2 = delete(5, tree2)
	printInOrder(tree2)
	fmt.Println("")
	tree2 = delete(4, tree2)
	printInOrder(tree2)
	fmt.Println("")
	tree2 = delete(3, tree2)
	printInOrder(tree2)
	fmt.Println("")
	tree2 = delete(2, tree2)
	printInOrder(tree2)
	fmt.Println("")
	tree2 = delete(1, tree2)
	fmt.Println("empty")
	//tree2 = *delete(1, &tree2)
	printInOrder(tree2)
	//fmt.Println("not empty")
}
