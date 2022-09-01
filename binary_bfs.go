package main

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

func main() {
	//arr := []int{1, -1, 2, 4, 3}
	//res := bfs(arr)
	//fmt.Println(res)
}

func bfs(root *TreeNode) {
	q := make([]*TreeNode, 1, 0)
	q[0] = root
	for len(q) > 0 {
		p := make([]*TreeNode, 0)
		for len(q) > 0 {
			temp := q[0]
			q = q[1:]
			fmt.Println(temp.value)
			if temp.left != nil {
				p = append(p, temp.left)
			}
			if temp.right != nil {
				p = append(p, temp.right)
			}
		}
		q = p
	}
}
