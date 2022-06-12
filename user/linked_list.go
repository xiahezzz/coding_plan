package user

type Node struct {
	Data interface{}
	Next *Node
}

type DoubleNode struct {
	Data interface{}
	Next *DoubleNode
	Pre  *DoubleNode
}

type NodeQueue struct {
	Head   *Node
	Tail   *Node
	length int
}

type NodeStack struct {
	Head   *Node
	length int
}
