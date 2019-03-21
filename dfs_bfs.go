package main

import (
	"math/rand"
	"fmt"
	"time"
)

type Node struct {
	left  *Node
	right *Node
	value int
}

func (n *Node) print() string {
	if n == nil {
		return "node is nil"
	}
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

func searchBFS(root *Node, target int) *Node {
	if root.value == target {
		return root
	}
	if root.left != nil && root.left.value == target {
		return root.left
	}

	if root.right != nil && root.right.value == target {
		return root.right
	}

	if root.left != nil {
		return searchBFS(root.left, target)
	}
	if root.right != nil {
		return searchBFS(root.right, target)
	}
	return nil
}

func main() {
	rand.Seed(time.Now().Unix())
	root := new(Node)
	tRoot := root
	for i := 0; i < 100000; i++ {
		val := rand.Intn(100000)
		tRoot.value = val
		if val%2 == 0 {
			tRoot.left = new(Node)
			tRoot = tRoot.left
		} else if val%9 == 0 {
			val = 55
		} else {
			tRoot.right = new(Node)
			tRoot = tRoot.right
		}
	}

	fmt.Println(searchBFS(root, 55).print())
}
