package medium

import (
	"bytes"
	"container/list"
	"fmt"
	"sort"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
LC 763. 划分字母区间
*/
func PartitionLabels(s string) []int {
	temp := map[byte]int{}
	for i := 0; i < len(s); i++ {
		temp[s[i]] = i
	}
	result := make([]int, 0)
	left := 0
	right := 0
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for j := 0; j < len(s); j++ {
		right = max(right, temp[s[j]])
		if right == j {
			result = append(result, right-left+1)
			left = right + 1
		}
	}
	return result
}

/*
LC 1604. 警告一小时内使用相同员工卡大于等于三次的人
*/
func AlertNames(keyName []string, keyTime []string) (ans []string) {
	timeMap := map[string][]int{}
	for i, name := range keyName {
		t := keyTime[i]
		hour := int(t[0]-'0')*10 + int(t[1]-'0')
		minute := int(t[3]-'0')*10 + int(t[4]-'0')
		timeMap[name] = append(timeMap[name], hour*60+minute)
	}
	for name, times := range timeMap {
		sort.Ints(times)
		for i, t := range times[2:] {
			if t-times[i] <= 60 {
				ans = append(ans, name)
				break
			}
		}
	}
	sort.Strings(ans)
	return
}

/*
LC 738. 单调递增的数字
*/
func MonotoneIncreasingDigits(n int) int {
	str := []rune(strconv.Itoa(n))
	flag := len(str)
	for i := len(str) - 1; i > 0; i-- {
		if str[i-1] > str[i] {
			flag = i
			str[i-1]--
		}
	}
	ans := 0
	for i := 0; i < len(str); i++ {
		if i >= flag {
			str[i] = '9'
		}
		ans = ans*10 + int(str[i]-'0')
	}
	return ans
}

/*
LC 6. N字型变换
*/
func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([][]string, numRows)
	index := 0
	row := 0
	s_len := len(s)
	for index < s_len {
		for index < s_len && row < numRows {
			ch := s[index]
			res[row] = append(res[row], string(ch))
			row++
			index++
		}
		if index == s_len {
			break
		}
		row = numRows - 2
		for index < s_len && row >= 0 {
			ch := s[index]
			res[row] = append(res[row], string(ch))
			row--
			index++
		}
		row += 2
	}
	ans := ""
	for _, v := range res {
		for _, v1 := range v {
			ans += v1
		}
	}
	return ans
}

/*
LC 739. 每日温度
*/
func DailyTemperatures(temperatures []int) []int {
	stack := list.New()
	t_len := len(temperatures)
	result := make([]int, t_len)
	stack.PushFront(0)
	for i := 1; i < t_len; i++ {
		// 小于栈头时插入
		if temperatures[i] < temperatures[stack.Front().Value.(int)] {
			stack.PushFront(i)
			// 等于栈头时插入
		} else if temperatures[i] == temperatures[stack.Front().Value.(int)] {
			stack.PushFront(i)
			// 大于栈头时循环弹出
		} else {
			for stack.Len() != 0 && temperatures[i] > temperatures[stack.Front().Value.(int)] {
				result[stack.Front().Value.(int)] = i - stack.Front().Value.(int)
				stack.Remove(stack.Front())
			}
			// 插入下标
			stack.PushFront(i)
		}
	}
	return result
}

/*
LC 12.整数转罗马数字
*/
func IntToRoman(num int) string {
	//罗马数字对应的阿拉伯数字
	symbols := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	ans := ""
	for _, v := range symbols {
		for num >= v.value {
			num -= v.value
			ans += v.symbol
		}
		if num == 0 {
			break
		}
	}
	return ans
}

/*
LC 22.括号生成(深度优先遍历)
*/
func GenerateParenthesis(n int) []string {
	ans := []string{}
	dfs := func(curStr string, left, right int, ans []string) []string { return []string{} }
	dfs = func(curStr string, left, right int, ans []string) []string {
		if left == 0 && right == 0 {
			ans = append(ans, curStr)
			return ans
		}
		if left > right {
			return ans
		}
		if left > 0 {
			ans = dfs(curStr+"(", left-1, right, ans)
		}
		if right > 0 {
			ans = dfs(curStr+")", left, right-1, ans)
		}
		return ans
	}
	ans = dfs("", n, n, ans)
	return ans
}

/*
LC 18.四数之和
*/
func FourSum(nums []int, target int) [][]int {
	//暂时不写
	return [][]int{}
}

/*
LC 24.两两交换链表中的节点
*/
func SwapPairs(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	node := pre
	for node.Next != nil && node.Next.Next != nil {
		start := node.Next
		end := node.Next.Next
		node.Next = end
		start.Next = end.Next
		end.Next = start
		node = start
	}
	return pre.Next
}

/*
LC 29.两数相除
*/
func Divide(dividend int, divisor int) int {
	sign := 1
	if (dividend ^ divisor) < 0 {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	ans := 0
	for i := 31; i >= 0; i-- {
		for dividend > 0 && (dividend>>i)-divisor >= 0 {
			dividend -= (divisor << i)
			ans += (1 << i)
		}
	}
	return sign * ans
}

/*
LC 31. 下一个排列
*/
func NextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	reverse := func(num []int) {
		for i, n := 0, len(num); i < n/2; i++ {
			num[i], num[n-1-i] = num[n-1-i], num[i]
		}
	}
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

/*
面试题 08.07. 无重复字符串的排列组合
*/
func Permutation(S string) []string {
	S_len := len(S)
	ans := []string{}
	used := make([]bool, S_len)
	dfs := func(path bytes.Buffer) {}
	dfs = func(path bytes.Buffer) {
		if path.Len() == len(S) {
			ans = append(ans, path.String())
			return
		}
		for i := 0; i < S_len; i++ {
			// 判断是否被使用
			if !used[i] {
				path.WriteByte(S[i])
				used[i] = true
				dfs(path)
				used[i] = false
				path.Truncate(path.Len() - 1)
				fmt.Printf("path.String(): %v\n", path.String())
			}
		}
	}
	dfs(bytes.Buffer{})
	return ans
}

/*
LC 34. 在排序数组中查找元素的第一个和最后一个位置
*/
func SearchRange(nums []int, target int) []int {
	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := sort.SearchInts(nums, target+1) - 1
	return []int{left, right}
}

/*
LC 38. 外观数列
*/
func CountAndSay(n int) string {
	n_map := make(map[int]string, n)
	n_map[1] = "1"
	// 1 - n 记录 对应的string
	for i := 2; i <= n; i++ {
		temp := ""
		s := n_map[i-1]
		// 循环删除前面相同的字符并记录出现次数
		for s != "" {
			i := 0
			top := s[0]
			for i+1 < len(s) && s[i+1] == top {
				i++
			}
			temp += fmt.Sprintf("%d%d", i+1, top-'0')
			s = s[i+1:]
		}
		n_map[i] = temp

	}
	return n_map[n]
}
