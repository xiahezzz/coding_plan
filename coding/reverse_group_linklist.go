package coding

import (
	"CodingPlan/user"
)

func GetKGroup(head *user.Node, k int) *user.Node {
	for ; k != 1 && head != nil; k-- {
		head = head.Next
	}
	return head
}

func ReverseNode(head *user.Node, tail *user.Node) *user.Node {
	var pre *user.Node
	var next *user.Node

	for head != tail {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}

	return pre
}

func ReverseByGroup(list *user.Node, k int) *user.Node {
	end := GetKGroup(list, k)
	if end == nil {
		return list
	}
	var lastEnd *user.Node
	var nextStart *user.Node

	head := end

	start := list
	nextStart = end.Next
	ReverseNode(start, nextStart)
	lastEnd = start

	for end != nil {
		start = nextStart
		end = GetKGroup(nextStart, k)
		if end == nil {
			lastEnd.Next = nextStart
			return head
		}
		nextStart = end.Next
		ReverseNode(start, nextStart)

		lastEnd.Next = end
		lastEnd = start
	}

	return head
}
