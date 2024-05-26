// i want to make a tree data strucuture here

package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {
	root := &Node{value: 1}
	root.left = &Node{value: 2}
	root.right = &Node{value: 3}
	root.left.left = &Node{value: 4}
	root.left.right = &Node{value: 5}
	root.right.left = &Node{value: 6}
	root.right.right = &Node{value: 7}

	fmt.Println(root)
}

// Output: &{1 0xc00000c030 0xc00000c060}
