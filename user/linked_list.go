package user

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type DoubleNode struct {
	Data int
	Next *DoubleNode
	Pre  *DoubleNode
}

func (n *Node) PrintList() {
	for n != nil {
		fmt.Printf("n.Data: %v\n", n.Data)
		n = n.Next
	}
}

func (n *Node) ReverseList() *Node {
	var pre *Node
	var next *Node

	for n != nil {
		next = n.Next
		n.Next = pre
		pre = n
		n = next
	}
	return pre
}

func RreversrList(head Node) *Node {
	cur := &head
	var pre *Node = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre //这句话最重要
	}
	return pre
}

func (n *Node) Update(num int) {
	temp := n.Next.Data
	n = n.Next
	n.Data = temp
}
