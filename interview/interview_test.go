package interview

import (
	"fmt"
	"testing"
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
