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

func TestMaxSlidingWindow(T *testing.T) {
	datas := []struct {
		nums []int
		k    int
	}{
		// {[]int{7, 2, 4}, 2},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
	}
	for k, data := range datas {
		result := MaxSlidingWindow(data.nums, data.k)
		fmt.Printf("第%d次测试的结果 : 结果数组：%v\n", k+1, result)
	}
}
