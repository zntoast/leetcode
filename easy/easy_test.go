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

func TestMerge(t *testing.T) {
	date := []struct {
		num1 []int
		m    int
		num2 []int
		n    int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
	}
	for _, data := range date {
		Merge(data.num1, data.m, data.num2, data.n)
		fmt.Printf("data.num1: %v\n", data.num1)
	}
}

func TestMaxDepth(t *testing.T) {
	right3 := &TreeNode{7, nil, nil}
	right2 := &TreeNode{15, nil, nil}
	right1 := &TreeNode{20, right2, right3}
	left1 := &TreeNode{9, nil, nil}
	root := &TreeNode{3, left1, right1}
	i := MaxDepth(root)
	fmt.Printf("i: %v\n", i)
}

func TestSortedArrayToBST(t *testing.T) {
	date := []struct {
		num1 []int
	}{
		{[]int{-10, -3, 0, 5, 9}},
	}
	for _, data := range date {
		SortedArrayToBST(data.num1)
	}
}

func TestGenerate(t *testing.T) {
	date := []struct {
		num int
	}{
		{5},
	}
	for k, data := range date {
		result := Generate(data.num)
		fmt.Printf("第%d次测试的结果 : 输出数组：%v\n", k+1, result)
	}
}

func TestConvertToTitle(t *testing.T) {
	date := []struct {
		num int
	}{
		{5},
		{26},
		// {701},
		//{2147483647},
	}
	for k, data := range date {
		result := ConvertToTitle(data.num)
		fmt.Printf("第%d次测试的结果 输入:%v 字符串：%v\n", k+1, data.num, result)
	}
}
