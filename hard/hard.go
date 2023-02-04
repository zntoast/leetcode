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

/*
LC 2071 你可以安排的最多任务数目
*/
func MaxTaskAssign(tasks []int, workers []int, pills int, strength int) int {
	k := min(len(tasks), len(workers))
	sort.Ints(tasks)
	sort.Ints(workers)
	check := func(n int) bool {
		p, wq, j := pills, make([]int, 0), len(workers)-1
		for i := n - 1; i >= 0; i-- {
			for ; j >= 0 && workers[j]+strength >= tasks[i]; j-- {
				wq = append(wq, workers[j])
			}
			if len(wq) == 0 {
				return false
			}
			if wq[0] >= tasks[i] {
				wq = wq[1:]
			} else {
				if p > 0 {
					p--
					wq = wq[:len(wq)-1]
				} else {
					return false
				}
			}
		}
		return true
	}
	return sort.Search(k, func(i1 int) bool {
		i1++
		return !check(i1)
	})
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
