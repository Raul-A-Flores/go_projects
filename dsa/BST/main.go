package main

import "fmt"

// Node

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert

func (n *Node) Insert(k int) {

	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {

		if n.Left == nil {

			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}

	}

}

// Search will take in a key value and return true if there is a node with that value

func (n *Node) Search(s int) bool {

	if n == nil {
		return false
	}
	if n.Key < s {
		// move right

		return n.Right.Search(s)

	} else if n.Key > s {
		// move left
		return n.Left.Search(s)

	}

	return true
}

func main() {

	tree := &Node{Key: 100}
	tree.Insert(500)
	tree.Insert(134)

	fmt.Println(tree)

	fmt.Println(tree.Search(500))

}
