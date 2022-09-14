package linkedlist

import (
	"fmt"
	"testing"
)

func TestFindNodeValue(T *testing.T) {
	hd1 := ListNode{1, nil}
	hd2 := ListNode{2, &hd1}
	hd3 := ListNode{3, &hd2}
	hd4 := ListNode{4, &hd3}
	fmt.Println(FindNodeValue(&hd4, 0))
}

func TestInsertHeadNode(T *testing.T) {
	hd1 := ListNode{1, nil}
	hd2 := ListNode{2, &hd1}
	hd3 := ListNode{3, &hd2}
	hd4 := ListNode{4, &hd3}
	fmt.Println(InsertHeadNode(&hd4, 5))
}

func TestDeleteNodeValue(T *testing.T) {
	hd1 := ListNode{1, nil}
	hd2 := ListNode{2, &hd1}
	hd3 := ListNode{3, &hd2}
	hd4 := ListNode{4, &hd3}
	fmt.Println(DeleteNodeValue(&hd4, 2))
}
