package user

import "fmt"

func (n *Node) PrintList() *Node {
	list := n
	for n != nil {
		fmt.Printf("n.Data: %v\n", n.Data)
		n = n.Next
	}
	return list
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

func (n *Node) GetArrList() []int {
	result := make([]int, 0)
	for i := 0; n != nil; i++ {
		result = append(result, n.Data.(int))
		n = n.Next
	}
	return result
}

func Queue() *NodeQueue {
	return &NodeQueue{Head: nil, Tail: nil}
}

func (nq *NodeQueue) Push(data interface{}) bool {
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

func (nq *NodeQueue) Length() int {
	return nq.length
}

func (nq *NodeQueue) Peek() interface{} {
	if nq.length == 0 {
		return nil
	}
	return nq.Head.Data
}

func (nq *NodeQueue) Pop() interface{} {
	if nq.length == 0 {
		return nq.Peek()
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

func (deq *DoubleNodeQueue) Length() int {
	return deq.length
}

func (deq *DoubleNodeQueue) IsEmpty() bool {
	return deq.length == 0
}

func (deq *DoubleNodeQueue) Push(data interface{}, method string) bool {
	if method != "left" && method != "right" {
		return false
	}

	dn := &DoubleNode{
		Data: data,
		Next: nil,
		Pre:  nil,
	}

	if deq.length == 0 {
		deq.Head = dn
		deq.Tail = dn
	}

	if method == "left" {
		deq.Head.Pre = dn
		dn.Next = deq.Head
		deq.Head = dn
	} else {
		deq.Tail.Next = dn
		dn.Pre = deq.Tail
		deq.Tail = dn
	}

	deq.length++
	return true
}

func (deq *DoubleNodeQueue) Peek(method string) interface{} {
	if method != "left" && method != "right" {
		return false
	}
	if deq.length == 0 {
		return nil
	}

	if method == "left" {
		return deq.Head.Data
	}
	return deq.Tail.Data
}

func (deq *DoubleNodeQueue) Pop(method string) interface{} {
	if method != "left" && method != "right" {
		return false
	}

	if deq.length == 1 {
		data := deq.Head.Data
		deq.Head = nil
		deq.Tail = nil

		deq.length--
		return data
	}

	if method == "left" {
		deq.Head = deq.Head.Next
		deq.Head.Pre = nil
	} else {
		deq.Tail = deq.Tail.Pre
		deq.Tail.Next = nil
	}

	deq.length--

	return deq.Head.Data
}
