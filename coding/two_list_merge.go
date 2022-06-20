//两个有序链表合并
package coding

func TwoListMerge(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var l1 *ListNode
	var l2 *ListNode
	var l1last *ListNode

	var head *ListNode

	if list1.Val <= list2.Val {
		head = list1
		l1last = list1
		l1 = list1.Next
		l2 = list2
	} else {
		head = list2
		l1last = list2
		l1 = list2.Next
		l2 = list1
	}

	for l1 != nil && l2 != nil {
		if l2.Val < l1.Val {
			l1last.Next = l2
			l2 = l2.Next
			l1last = l1last.Next
			l1last.Next = l1
		} else {
			l1last = l1last.Next
			l1 = l1.Next
		}
	}
	if l2 != nil {
		l1last.Next = l2
	}
	return head
}
