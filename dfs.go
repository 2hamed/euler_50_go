package main

import (
	"math/rand"
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	value int
}

func (n *Node) print() string {
	var l = -10000
	if n.left != nil {
		l = n.left.value
	}

	var r = -10000
	if n.right != nil {
		r = n.right.value
	}

	return fmt.Sprintf("value: %d, left: %d, right: %d", n.value, l, r)
}

func searchDFS(root *Node, target int) *Node {
	if root.value == target {
		return root
	}
	fmt.Println(root.print())
	if root.left != nil {
		return searchDFS(root.left, target)
	}
	if root.right != nil {
		return searchDFS(root.right, target)
	}

	return nil
}

func main() {
	root := new(Node)
	tRoot := root
	for i := 0; i < 10; i++ {
		val := rand.Intn(1000)
		tRoot.value = val
		if val%2 == 0 {
			tRoot.left = new(Node)
			tRoot = tRoot.left
		} else {
			tRoot.right = new(Node)
			tRoot = tRoot.right
		}
	}
	tRoot.right = new(Node)
	tRoot.right.value = 55

	fmt.Println(searchDFS(root, 55).print())
}
