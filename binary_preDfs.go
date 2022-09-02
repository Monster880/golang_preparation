package main

import "fmt"

type TreeNode1 struct {
	left  *TreeNode1
	right *TreeNode1
	value int
}

//func main() {
//	root := &TreeNode1{value: 0}
//	root.left = &TreeNode1{value: 1}
//	root.right = &TreeNode1{value: 2}
//	dfs(root)
//}

func dfs(root *TreeNode1) {
	fmt.Println(root.value)
	if root.left != nil {
		dfs(root.left)
	}
	if root.right != nil {
		dfs(root.right)
	}
}
