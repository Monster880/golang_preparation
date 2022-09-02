package main

type LinkedListNode struct {
	next  *LinkedListNode
	value int
}

//func main() {
//	root := &LinkedListNode{value: 0}
//	dummyNode := root
//	dummyNode.next = &LinkedListNode{value: 1}
//	dummyNode = dummyNode.next
//	dummyNode.next = &LinkedListNode{value: 2}
//	res := reverseList(root)
//	for res != nil {
//		fmt.Println(res.value)
//		res = res.next
//	}
//}

func reverseList(root *LinkedListNode) *LinkedListNode {
	dummyNode := root
	var pre *LinkedListNode
	pre = nil
	for dummyNode != nil {
		next := dummyNode.next
		dummyNode.next = pre
		pre = dummyNode
		dummyNode = next
	}
	return pre
}
