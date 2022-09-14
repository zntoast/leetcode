package basic

import (
	"fmt"
	"testing"
)

func TestReversePrint(T *testing.T) {
	l2 := ListNode{2, nil}
	l3 := ListNode{3, &l2}
	l1 := ListNode{1, &l3}
	fmt.Println(ReversePrint(&l1))
}

func TestReverseLeftWords(T *testing.T) {
	s := "abcdefg"
	fmt.Println(ReverseLeftWords(s, 2))
}
