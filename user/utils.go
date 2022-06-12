package user

import "fmt"

func (n *Node) PrintList() {
	for n != nil {
		fmt.Printf("n.Data: %v\n", n.Data)
		n = n.Next
	}
}

func (n *Node) Length() int {
	length := 0
	for n != nil {
		length += 1
		n = n.Next
	}
	return length
}

func (n *DoubleNode) PrintList() {
	for n != nil {
		fmt.Printf("n.Data: %v\n", n.Data)
		n = n.Next
	}
}

func (n *Node) Update(num int) {
	temp := n.Next.Data
	n = n.Next
	n.Data = temp
}

func Queue() *NodeQueue {
	return &NodeQueue{Head: nil, Tail: nil}
}

func (nq *NodeQueue) EnterQueue(data interface{}) bool {
	n := &Node{
		Data: data,
		Next: nil,
	}
	if nq.Head == nil {
		nq.Head = n
		nq.Tail = n
		nq.length++
		return true
	}
	nq.Tail.Next = n
	nq.Tail = n
	nq.length++
	return true
}

func (nq *NodeQueue) LengthQueue() int {
	return nq.length
}

func (nq *NodeQueue) PeekQueue() interface{} {
	if nq.length == 0 {
		return nil
	}
	return nq.Head.Data
}

func (nq *NodeQueue) PopQueue() interface{} {
	if nq.LengthQueue() == 0 {
		return nq.PeekQueue()
	}
	temp := nq.Head
	nq.Head = nq.Head.Next
	nq.length--

	//如果不加这一条会产生脏数据，无法被回收。（虽然对实际逻辑没有产生影响）
	if nq.Head == nil {
		nq.Tail = nil
	}
	return temp.Data
}

func (nq *NodeQueue) IsEmpty() bool {
	return nq.length == 0
}

func ReverseNode(n *Node) *Node {
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

func ReverseDNode(n *DoubleNode) *DoubleNode {
	var pre *DoubleNode
	var next *DoubleNode

	for n != nil {
		next = n.Next

		n.Next = pre
		n.Pre = next

		pre = n
		n = next
	}

	return pre
}

func Stack() *NodeStack {
	return &NodeStack{}
}

func (ns *NodeStack) Length() int {
	return ns.length
}

func (ns *NodeStack) IsEmpty() bool {
	return ns.length == 0
}

func (ns *NodeStack) Push(data interface{}) bool {
	n := &Node{
		Data: data,
		Next: nil,
	}

	n.Next = ns.Head
	ns.Head = n
	ns.length++
	return true
}

func (ns *NodeStack) Peek() interface{} {
	if ns.length == 0 {
		return nil
	}

	return ns.Head.Data
}

func (ns *NodeStack) Pop() interface{} {
	if ns.length == 0 {
		return nil
	}
	temp := ns.Head
	ns.Head = ns.Head.Next
	ns.length--
	return temp
}
