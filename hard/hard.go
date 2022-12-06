package hard

import (
	"fmt"
	"sort"
)

/*
LC 805 数组的均值分割
*/
func SplitArraySameAverage(nums []int) bool {
	//排序
	sort.Ints(nums)
	return false
}

func depth() {
	depth := 0
	for i := 10; i > 0; i >>= 1 {
		depth++
	}
	fmt.Printf("depth: %v\n", depth)
}

type Str struct {
	Name string
}
