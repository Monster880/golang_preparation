package main

import "fmt"

type TreeNode2 struct {
	left  *TreeNode2
	right *TreeNode2
	value int
}

func main() {
	root := &TreeNode2{value: 0}
	root.left = &TreeNode2{value: 1}
	root.right = &TreeNode2{value: 2}
	stack := make([]*TreeNode2, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		for root.left != nil {
			fmt.Println(root.value)
			stack = append(stack, root.left)
			root = root.left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.right
		}
	}
}
