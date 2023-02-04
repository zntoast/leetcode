package hard

import (
	"fmt"
	"testing"
)

func TestTepth(t *testing.T) {
	depth()
}

func TestMaxTaskAssign(t *testing.T) {
	date := []struct {
		tasks    []int
		workers  []int
		pills    int
		strength int
	}{
		{[]int{3, 2, 1}, []int{0, 3, 3}, 1, 1},
		{[]int{5, 4}, []int{0, 0, 0}, 1, 5},
		{[]int{10, 15, 30}, []int{0, 10, 10, 10, 10}, 3, 10},
		{[]int{5, 9, 8, 5, 9}, []int{1, 6, 4, 2, 6}, 1, 5},
	}
	for k, d := range date {
		count := MaxTaskAssign(d.tasks, d.workers, d.pills, d.strength)
		fmt.Printf("第%d次测试的结果 :count: %v\n", k+1, count)
	}
}
