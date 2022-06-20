//两链表之和

package coding

type ListNode struct {
	Next *ListNode
	Val  int
}

func TwoListSum(list1 *ListNode, list2 *ListNode) {
	last := 0
	this := 0
	for {
		this = list1.Val + list2.Val + last
		list1.Val = this % 10
		last = this / 10
		if list1.Next == nil || list2.Next == nil {
			break
		}
		list1 = list1.Next
		list2 = list2.Next
	}

	if list2.Next != nil {
		list1.Next = list2.Next
	}

	for last != 0 {
		if list1.Next == nil {
			temp := &ListNode{
				Val:  1,
				Next: nil,
			}
			list1.Next = temp
			break
		}

		list1 = list1.Next
		this = list1.Val + last
		list1.Val = this % 10
		last = this / 10
	}
}
