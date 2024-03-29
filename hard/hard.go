package hard

import (
	"fmt"
	"math"
	"sort"
	"strings"
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

// 插入数据
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

// 获取数据并删除
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

// 串联所有单词的字串
func FindSubstring(s string, words []string) []int {
	result := []string{}
	w_len := len(words)
	used := map[int]bool{}
	for k := range words {
		used[k] = false
	}
	dfs := func(buff string, deep int) {}
	dfs = func(buff string, deep int) {
		if deep == w_len {
			result = append(result, buff)
			return
		}
		for i := 0; i < w_len; i++ {
			if !used[i] {
				used[i] = true
				buff += words[i]
				dfs(buff, deep+1)
				used[i] = false
				buff = buff[:len(buff)-len(words[0])]
			}
		}
	}
	dfs("", 0)
	ans := []int{}
	exist := map[int]struct{}{}
	for _, v := range result {
		indexs := findAllSubstringIndices(s, v)
		for _, index := range indexs {
			_, ok := exist[index]
			if index != -1 && !ok {
				ans = append(ans, index)
				exist[index] = struct{}{}
			}
		}
	}
	return ans
}

func findAllSubstringIndices(s, sub string) []int {
	indices := []int{}
	start := 0
	for {
		index := strings.Index(s[start:], sub)
		if index == -1 {
			break
		}
		indices = append(indices, start+index)
		start += index + 1
	}
	return indices
}

// LC 514. 自由之路
func FindRotateSteps(ring string, key string) int {
	ans := math.MaxInt32
	dfs := func(nums []byte, index, left, right, value int) {}
	dfs = func(nums []byte, index, left, right, value int) {
		if index == len(key) {
			ans = min(ans, value)
			return
		}
		for {
			// 逆时针
			if nums[left] == key[index] {
				//  我们只需要1步来拼写这个字符
				value++
				// 当前转盘顺序
				dfs(append(nums[left:], nums[:left]...), index+1, 0, 0, value)
			}
			// 顺时针
			if nums[right] == key[index] {
				//  我们只需要1步来拼写这个字符
				value++
				// 当前转盘顺序
				dfs(append(nums[right:], nums[:right]...), index+1, 0, 0, value)
			}
			// 移动转盘
			right = len(ring) - left - 1
			left++
			value++
			if right <= left {
				break
			}
		}
	}
	dfs([]byte(ring), 0, 0, 0, 0)
	return ans
}
