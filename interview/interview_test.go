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
	list := []int{5, 10, 2, 1, 6, 7, 4, 4, 3, 9, 5, 10, 2, 1, 6, 7, 4, 4, 3, 9, 5, 10, 2, 1, 6, 7, 4, 4, 3, 9, 5, 10, 2, 1, 6, 7, 4, 4, 3, 9, 5, 10, 2, 1, 6, 7, 4, 4, 3, 9}
	start := time.Now()
	// 快速排序
	// FastSort(list)
	// 冒泡排序
	BubbleSort(list)
	end := time.Now()
	fmt.Printf("时间开销:%v \n", end.Sub(start))
	fmt.Printf("list: %v\n", list)
}
