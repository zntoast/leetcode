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

func TestTitleToNumber(t *testing.T) {
	date := []struct {
		columnTitle string
	}{
		{"A"},
		{"AB"},
		{"ZY"},
	}
	for k, data := range date {
		result := TitleToNumber(data.columnTitle)
		fmt.Printf("第%d次测试的结果 输入:%v 字符串：%v\n", k+1, data.columnTitle, result)
	}
}

func TestHammingWeight(t *testing.T) {
	date := []struct {
		num uint32
	}{
		{00000000000000000000000000001011},
	}
	for k, data := range date {
		result := HammingWeight(data.num)
		fmt.Printf("第%d次测试的结果 输入:%v 存在：%v 个1\n", k+1, data.num, result)
	}
}

func TestContainsNearbyDuplicate(t *testing.T) {
	date := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 2, 3, 1}, 3},
		{[]int{1, 2, 3, 1, 2, 3}, 2},
	}
	for k, data := range date {
		result := ContainsNearbyDuplicate(data.nums, data.k)
		fmt.Printf("第%d次测试的结果 %v \n", k+1, result)
	}
}

func TestRemoveElements(t *testing.T) {
	node5 := &ListNode{5, nil}
	node4 := &ListNode{4, node5}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{6, node3}
	node1 := &ListNode{2, node2}
	head := &ListNode{1, node1}

	// node3 := &ListNode{7, nil}
	// node2 := &ListNode{7, node3}
	// node1 := &ListNode{7, node2}
	// head := &ListNode{7, node1}
	result := RemoveElements(head, 6)
	fmt.Printf("result: %v\n", result)
}

func TestMyStack(t *testing.T) {
	m := MyStack{}
	for i := 0; i < 5; i++ {
		m.Push(i)
	}
	fmt.Printf("插入数据m: %v\n", m)
	i := m.Pop()
	fmt.Printf("i: %v\n", i)
	fmt.Printf("剩余数据i: %v\n", m)
	i1 := m.Top()
	fmt.Printf("i1: %v\n", i1)
	fmt.Printf("剩余数据i1: %v\n", m)
}

func TestSummaryRanges(t *testing.T) {
	date := []struct {
		nums []int
	}{
		{[]int{0, 1, 2, 4, 5, 7}},
		{[]int{0, 2, 3, 4, 6, 8, 9}},
	}
	for k, data := range date {
		result := SummaryRanges(data.nums)
		fmt.Printf("第%d次测试的结果  result：%v 个\n", k+1, result)
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	date := []struct {
		num int
	}{
		{1},
		{16},
		{32},
		{64},
		{128},
		{256},
		{512},
		{1056},
	}
	for k, data := range date {
		result := IsPowerOfTwo(data.num)
		fmt.Printf("第%d次测试的结果  输入:%v 输入%v是2的ning \n", k+1, data.num, result)
	}
}

func testadddigits(t *testing.T) {
	date := []struct {
		num int
	}{
		{1},
		{0},
		{100},
	}
	for k, data := range date {
		result := AddDigits(data.num)
		fmt.Printf("第%d次测试的结果  输入:%v 输出:%v\n", k+1, data.num, result)
	}
}

func TestIsUgly(t *testing.T) {
	date := []struct {
		num int
	}{
		{6},
		{1},
		{14},
	}
	for k, data := range date {
		result := IsUgly(data.num)
		fmt.Printf("第%d次测试的结果  输入:%v IS丑数:%v\n", k+1, data.num, result)
	}
}

func TestMissingNumber(t *testing.T) {
	date := []struct {
		nums []int
	}{
		{[]int{3, 0, 1}},
		{[]int{0, 1}},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}},
	}
	for k, data := range date {
		result := MissingNumber(data.nums)
		fmt.Printf("第%d次测试的结果  输入:%v value:%v\n", k+1, data.nums, result)
	}
}

func TestWordPattern(t *testing.T) {
	date := []struct {
		pattern string
		s       string
	}{
		{"abba", "dog cat cat dog"},
		{"abba", "dog cat cat fish"},
		{"aaaa", "dog cat cat dog"},
	}
	for k, data := range date {
		result := WordPattern(data.pattern, data.s)
		fmt.Printf("第%d次测试的结果  value:%v\n", k+1, result)
	}
}

func TestCountBits(t *testing.T) {
	date := []struct {
		n int
	}{
		{5},
	}
	for k, data := range date {
		result := CountBits(data.n)
		fmt.Printf("第%d次测试的结果  value:%v\n", k+1, result)
	}
}

func TestCount(t *testing.T) {
	date := []struct {
		n int
	}{
		{5},
		{3},
		{2},
		{10},
		{16},
	}
	for k, data := range date {
		result := Count(data.n)
		fmt.Printf("第%d次测试的结果  value:%v\n", k+1, result)
	}
}

func TestGuessNumber(t *testing.T) {
	date := []struct {
		n int
	}{
		{10},
	}
	for k, data := range date {
		result := GuessNumber(data.n)
		fmt.Printf("第%d次测试的结果  value:%v\n", k+1, result)
	}
}

func TestDistributeCandies(t *testing.T) {
	date := []struct {
		candyType []int
	}{
		{[]int{1, 1, 2, 2, 3, 3}},
		{[]int{1, 1, 2, 3}},
		{[]int{6, 6, 6, 6}},
	}
	for k, data := range date {
		result := DistributeCandies(data.candyType)
		fmt.Printf("第%d次测试的结果  value:%v\n", k+1, result)
	}
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	date := []struct {
		nums []int
	}{
		{[]int{1, 1, 0, 1, 1, 1}},
		{[]int{1, 1, 0, 1, 0, 0}},
		{[]int{1, 1, 1, 1}},
	}
	for k, data := range date {
		result := FindMaxConsecutiveOnes(data.nums)
		fmt.Printf("第%d次测试的结果  数组:%v value:%v\n", k+1, data.nums, result)
	}
}

func TestConstructRectangle(t *testing.T) {
	date := []struct {
		area int
	}{
		{4},
		{37},
	}
	for k, data := range date {
		result := ConstructRectangle(data.area)
		fmt.Printf("第%d次测试的结果  数组:%v value:%v\n", k+1, data.area, result)
	}
}
