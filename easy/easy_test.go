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
