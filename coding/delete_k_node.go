package coding

func GetLength(head *ListNode) int {
	result := 0
	for head != nil {
		result++
		head = head.Next
	}
	return result
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	length := GetLength(head)
	if length == n {
		return head.Next
	}
	result := head

	for i := 0; i < length-n-1; i++ {
		head = head.Next
	}

	pre := head
	head = head.Next
	if head.Next == nil {
		pre.Next = nil
		return result
	}
	pre.Next = head.Next
	head.Next = nil
	return result
}

func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	var p1 *ListNode
	var p2 *ListNode

	p2 = head
	p1 = head
	for ; n != -1; n-- {
		p2 = p2.Next
	}
	if p2 == nil {
		return head.Next
	}
	for p2 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	p1.Next = p1.Next.Next
	return head
}
