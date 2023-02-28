package interview

import (
	"fmt"
	"testing"
	"time"
)

func TestFlower(T *testing.T) {
	list := []int{0, 0, 0, 0, 1, 0}
	n := 1
	fmt.Println(Flower(list, n))
}

func TestFindString(T *testing.T) {
	str1 := "abcccccccfg"
	str2 := "bcccccfg"
	fmt.Println(FindString(str1, str2))
}

func TestFibonacciSequence(t *testing.T) {
	FibonacciSequence(1000)
}

func TestFastSort(t *testing.T) {
	list := []int{5, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7,
		5, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 44, 5, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7,
		5, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 44, 5, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7,
		5, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 44, 5, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7,
		5, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 45, 10, 2, 1, 6, 7, 44}
	start := time.Now()
	// 快速排序
	// FastSort(list)
	// 冒泡排序
	// BubbleSort(list)
	// 插入排序
	// InsertSort(list)
	// 分治排序
	MergeSort(list, 0, len(list)-1)
	end := time.Now()
	fmt.Printf("时间开销:%v \n", end.Sub(start))
	fmt.Printf("list: %v\n", list)
}

func TestEratosthenes(t *testing.T) {
	date := []int{15, 98, 100, 1000, 10000}
	for k, v := range date {
		c := Eratosthenes(v)
		fmt.Printf("第 %d 次测试 有 %d 个素数 \n", k+1, c)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	date := []struct {
		nums []int
	}{
		{[]int{1, 1, 2, 3, 3, 4, 5}},
		{[]int{2, 2, 2, 3, 3, 4, 5}},
		{[]int{2, 2, 3, 3, 4, 5}},
		{[]int{2, 2, 2, 2, 2, 3, 3, 4, 5}},
		{[]int{2, 2, 2, 3, 3, 4, 4, 5, 5}},
	}
	for k, v := range date {
		num := RemoveDuplicates(v.nums)
		fmt.Printf("第 %d 次测试 数组新长度: %d \n", k+1, num)
	}
}

func TestPivotIndex(t *testing.T) {
	date := []struct {
		nums []int
	}{
		{[]int{1, 7, 3, 6, 5, 6}},
		{[]int{1, 7, 3, 6, 5, 6, 2}},
		{[]int{1, 7, 3, 6, 2, 5, 6}},
		{[]int{1, 7, 3, 6, 5, 5, 6}},
	}
	for k, v := range date {
		index := PivotIndex(v.nums)
		fmt.Printf("第 %d 次测试 中间下标: %d \n", k+1, index)
	}
}

func TestBinarySearch(t *testing.T) {
	date := []struct {
		num int
	}{
		{24},
		{28},
		{29},
		{88},
		{100},
		{105},
	}
	for k, v := range date {
		index := BinarySearch(v.num)
		fmt.Printf("第 %d 次测试 %d平方根: %d \n", k+1, v.num, index)
	}
}

func TestRecursion(t *testing.T) {
	h4 := &ListNote{4, nil}
	h3 := &ListNote{3, h4}
	h2 := &ListNote{2, h3}
	head := &ListNote{1, h2}
	new_head := Recursion(head)
	fmt.Printf("new_head: %v\n", new_head)
}

func TestSort(t *testing.T) {
	date := []struct {
		nums []int
	}{
		{[]int{-4, -6, 0, 2, 5}},
		{[]int{-1, -4, 0, 2, 5}},
		{[]int{-4, -6, 0, 2, 5, 10}},
	}
	for k, v := range date {
		result := Sort(v.nums)
		fmt.Printf("第 %d 次测试 数组的最大乘积:%d \n", k+1, result)
	}
}

func TestSolution(t *testing.T) {
	date := []struct {
		nums  []int
		tager int
	}{
		{[]int{-4, -6, 0, 2, 5}, 28},
		{[]int{-1, -4, 0, 2, 5}, -5},
		{[]int{-4, -6, 0, 2, 5, 10}, 10},
		{[]int{-4, -6, 0, 2, 2, 5, 9, 10}, 10},
		{[]int{-4, -6, 0, 5, 10}, -10},
		{[]int{-4, -6, 0, 2, 10}, 10},
	}
	for k, v := range date {
		result := Solution(v.nums, v.tager)
		fmt.Printf("第 %d 次测试 数组下标:%d \n", k+1, result)
	}
}

func TestMerge(t *testing.T) {
	date := []struct {
		num1 []int
		num2 []int
	}{
		// {[]int{1, 5, 9, 10, 10}, []int{2, 4, 8, 9}},
		// {[]int{1, 3, 4, 5, 9, 10, 10}, []int{2, 4, 5, 6, 8, 9}},
		{[]int{1, 5, 6, 7, 8}, []int{2, 4, 8, 9}},
	}
	for k, v := range date {

		v.num1 = append(v.num1, v.num2...)
		//仅限两个数组长度相差不超过1
		Merge(v.num1, 0, len(v.num1)-1)
		fmt.Printf("第 %d 次测试 数组排序后:%v \n", k+1, v.num1)
	}
}

func TestArrangeCoins(t *testing.T) {
	date := []struct {
		tager int
	}{
		{16},
		{18},
		{90},
		{46},
		{55},
		{464},
	}
	for k, v := range date {
		result, count := ArrangeCoins(v.tager)
		fmt.Printf("第 %d 次测试 阶梯的行数:%d ,累加:%d \n", k+1, result, count)
	}
}

func TestFindStrings(t *testing.T) {
	date := []struct {
		a string
		b string
	}{
		{"服务提供商提供不了服务服务", "服务"},
		{"aadbbdowabdow", "do"},
		{"aadbbdowabdow", "dow"},
	}
	for k, v := range date {
		count := FindStrings(v.a, v.b)
		fmt.Printf("第 %d 次测试 找到:%d个 \n", k+1, count)
	}
}

func TestStrToInt(t *testing.T) {
	date := []struct {
		a string
	}{
		{" "},
		// {"3.14"},
		{"  -42"},
		{"42"},
	}
	for k, v := range date {
		reuslt := StrToInt(v.a)
		fmt.Printf("第 %d 次测试 值:%d \n", k+1, reuslt)
	}
}
