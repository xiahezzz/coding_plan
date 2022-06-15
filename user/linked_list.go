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

type DoubleNodeQueue struct {
	Head   *DoubleNode
	Tail   *DoubleNode
	length int
}

type NodeUser interface {
	Length() int
	IsEmpty() bool
	Push(data interface{}) bool
	Peek() interface{}
	Pop() interface{}
}

type DoubleNodeUser interface {
	Length() int
	IsEmpty() bool
	Push(data interface{}, method string) bool
	Peek(method string) interface{}
	Pop(method string) interface{}
}
