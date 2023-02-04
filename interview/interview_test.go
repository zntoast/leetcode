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
		fmt.Printf("第 %d 次测试 平方根: %d \n", k+1, index)
	}
}
