package easy

import (
	"fmt"
	"testing"
)

func TestLargestSumAfterKNegations(t *testing.T) {
	date := []struct {
		nums []int
		k    int
	}{
		{[]int{4, 2, 3}, 1},
		{[]int{3, -1, 0, 2}, 3},
		{[]int{2, -3, -1, 5, -4}, 2},
	}
	for k, data := range date {
		result := LargestSumAfterKNegations(data.nums, data.k)
		fmt.Printf("第%d次测试的结果 : 数组合：%v\n", k+1, result)
	}
}

func TestRomanToInt(t *testing.T) {
	date := []struct {
		s string
	}{
		{"M"},
		{"MCMXCIV"},
	}
	for k, data := range date {
		result := RomanToInt(data.s)
		fmt.Printf("第%d次测试的结果 : 数组合：%v\n", k+1, result)
	}
}

func TestSearchInsert(t *testing.T) {
	date := []struct {
		nums   []int
		target int
	}{
		{[]int{1, 3, 5, 6}, 5},
		{[]int{1, 3, 5, 6}, 2},
		{[]int{1, 3, 5, 6}, 7},
	}
	for k, data := range date {
		result := SearchInsert(data.nums, data.target)
		fmt.Printf("第%d次测试的结果 : 插入的下标为：%v\n", k+1, result)
	}
}

func TestMySqrt(t *testing.T) {
	date := []struct {
		num int
	}{
		//{4},
		//{8},
		{1},
	}
	for k, data := range date {
		result := MySqrt(data.num)
		fmt.Printf("第%d次测试的结果 : 插入的下标为：%v\n", k+1, result)
	}
}

func TestDeleteDuplicates(t *testing.T) {
	n2 := &ListNode{2, nil}
	n1 := &ListNode{1, n2}
	head := &ListNode{1, n1}
	newHead := DeleteDuplicates(head)
	node := newHead
	for node != nil {
		fmt.Printf("node.Val: %v\n", node.Val)
		node = node.Next
	}
}
