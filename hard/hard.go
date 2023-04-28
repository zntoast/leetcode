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

/*
LC 42. 接雨水
*/
func Trap(height []int) int {
	stack := []int{}
	ans := 0
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			height := min(height[left], h) - height[top]
			width := i - left - 1
			ans += height * width
		}
		stack = append(stack, i)
	}
	return ans
}

/*
LC 1172.餐盘栈
*/
type DinnerPlates struct {
	Data     [][]int
	Capacity int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		Data:     make([][]int, 0),
		Capacity: capacity,
	}
}

func (this *DinnerPlates) Push(val int) {
	if len(this.Data) == 0 {
		this.Data = append(this.Data, []int{val})
		return
	}
	for i := 0; i < len(this.Data); i++ {
		// 未装满时
		if len(this.Data[i]) < this.Capacity {
			this.Data[i] = append(this.Data[i], val)
			return
			// 装满并且是最后一个
		} else if len(this.Data[i]) == this.Capacity && i == len(this.Data)-1 {
			this.Data = append(this.Data, []int{val})
			return
			// 装满
		} else {
			continue
		}
	}
}

func (this *DinnerPlates) Pop() int {
	ans := -1
	if len(this.Data) == 0 {
		return ans
	}
	temp := this.Data[len(this.Data)-1]
	temp_len := len(temp)
	ans = temp[temp_len-1]
	this.Data[len(this.Data)-1] = this.Data[len(this.Data)-1][:temp_len-1]
	if len(this.Data[len(this.Data)-1]) == 0 {
		this.Data = this.Data[:len(this.Data)-1]
	}
	return ans
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if index < 0 || index >= len(this.Data) {
		return -1
	}
	temp := this.Data[index]
	temp_len := len(temp)
	if temp_len == 0 {
		return -1
	}
	ans := temp[temp_len-1]
	this.Data[index] = this.Data[index][:temp_len-1]
	return ans
}
