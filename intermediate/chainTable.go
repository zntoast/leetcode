package intermediate

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
LC 两数相加
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var left *ListNode
	var right *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		value := carry
		if l1 != nil {
			value += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			value += l2.Val
			l2 = l2.Next
		}
		carry = value / 10
		value = value % 10
		if left == nil {
			left = &ListNode{value, nil}
			right = left
		} else {
			right.Next = &ListNode{value, nil}
			right = right.Next
		}
	}
	if carry > 0 {
		right.Next = &ListNode{carry, nil}
	}
	return left
}

/*
LC 奇偶链表
*/
func OddEvenList(head *ListNode) *ListNode {
	var left *ListNode
	temp := new(ListNode)
	var right *ListNode
	temp1 := new(ListNode)
	node := head
	count := 1
	for node != nil {
		if count%2 == 1 {
			if left == nil {
				temp = &ListNode{node.Val, nil}
				left = temp
			} else {
				temp.Next = &ListNode{node.Val, nil}
				temp = temp.Next
			}
		}
		if count%2 == 0 {
			if right == nil {
				temp1 = &ListNode{node.Val, nil}
				right = temp1
			} else {
				temp1.Next = &ListNode{node.Val, nil}
				temp1 = temp1.Next
			}
		}
		node = node.Next
		count++
	}
	temp.Next = right
	return left
}

/*
LC 相交链表
*/
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	set := make(map[*ListNode]int)
	for headA != nil {
		set[headA]++
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := set[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
