package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

//func main() {
//	root := &ListNode{value: 0}
//	dummyNode := root
//	root.next = &ListNode{value: 1}
//	root = root.next
//	root.next = &ListNode{value: 2, next: nil}
//	reverseListNode(dummyNode)
//}

func reverseListNode(root *ListNode) *ListNode {
	dummyNode := root
	preNode := dummyNode
	for dummyNode != nil {
		next := dummyNode.next
		dummyNode.next = preNode
		preNode = dummyNode
		dummyNode = next
	}
	//for preNode != nil {
	//	fmt.Println(preNode)
	//	preNode = preNode.next
	//}
	fmt.Println(preNode)
	preNode = preNode.next
	fmt.Println(preNode)
	preNode = preNode.next
	fmt.Println(preNode)
	return preNode
}
