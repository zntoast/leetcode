package medium

import (
	"container/list"
	"sort"
	"strconv"
)

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
