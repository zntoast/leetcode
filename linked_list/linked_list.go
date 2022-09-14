package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

//删除链表头节点
func DeleteHeadNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

//删除链表的倒数第N个节点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	leftNode := head
	rightNode := head
	for i := 0; i < n; i++ {
		rightNode = rightNode.Next
	}
	if rightNode == nil {
		return head.Next
	}
	for rightNode.Next != nil {
		leftNode = leftNode.Next
		rightNode = rightNode.Next
	}
	leftNode.Next = leftNode.Next.Next
	return head
}

//反转链表
func ReverseList(head *ListNode) *ListNode {
	var newHead *ListNode = nil
	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}
	return newHead
}

//查找第n个节点的值
func FindNodeValue(node *ListNode, n int) int {
	head := node
	for n > 0 {
		if head.Next == nil {
			return 0
		}
		head = head.Next
		n--
	}
	return head.Val
}

//删除给定值的节点
func DeleteNodeValue(node *ListNode, target int) *ListNode {
	head := node
	for head != nil && head.Next != nil {
		if head.Next.Val == target {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return node
}

//头部插入节点
func InsertHeadNode(node *ListNode, a int) *ListNode {
	// dummyhead := new(ListNode)
	// dummyhead.Next = node
	head := new(ListNode)
	head.Val = a
	head.Next = node
	return head
}
