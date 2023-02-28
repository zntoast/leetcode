package shell

import (
	"fmt"
	"testing"
)

func TestArrayRankTransform(T *testing.T) {
	datas := []struct {
		arr []int
	}{
		{[]int{40, 10, 20, 30}},
		{[]int{37, 12, 28, 9, 100, 56, 80, 5, 12}},
	}
	for _, v := range datas {
		fmt.Printf("ArrayRankTransform(v.arr): %v\n", ArrayRankTransform(v.arr))
	}
}
