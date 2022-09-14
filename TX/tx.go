package tx

import (
	"math"
	"sort"
	"strconv"
	"unicode"
)

//两数相加 (完成)
//寻找两个正序数组的中位数(完成)
//最长回文子串
func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := center(s, i, i)
		left2, right2 := center(s, i, i+1)
		if right1-left1 > end-start {
			end, start = right1, left1
		}
		if right2-left2 > end-start {
			end, start = right2, left2
		}
	}
	return s[start:end]
}
func center(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left, right = left-1, right+1
	}
	return left + 1, right - 1
}

//7.整数反转 (完成)
func Reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	newX := 0
	for x > 0 {
		yu := x % 10
		newX = newX*10 + yu
		x = x / 10
	}
	newX = sign * newX
	if newX < math.MinInt32 || newX > math.MaxInt32 {
		return 0
	}
	return newX
}

//8字符串转整数
func MyAtoi(s string) int {
	//去除前导空格
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			s = s[i:]
			break
		}
	}
	if len(s) == 0 {
		return 0
	}
	sign := 1
	if string(s[0]) == "-" && len(s) >= 2 {
		sign = -sign
		s = s[1:]
	} else if string(s[0]) == "+" && len(s) >= 2 {
		s = s[1:]
	}
	var s1 string
	for _, v := range s {
		if unicode.IsNumber(v) {
			s1 += string(v)
		} else {
			break
		}
	}
	num, _ := strconv.Atoi(s1)
	if sign*num < math.MinInt32 {
		return math.MinInt32
	} else if sign*num > math.MaxInt32 {
		return math.MaxInt32
	}
	return sign * num
}

//9.回文数
func IsPalindrome(x int) bool {
	aa := x
	if x < 0 {
		return false
	}
	newNum := 0
	for x > 0 {
		y := x % 10
		x = x / 10
		newNum = newNum*10 + y
	}
	return newNum == aa
}

//14.最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

//15.三数之和
func ThreeSum(nums []int) [][]int {
	anx := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {
		target := -1 * nums[first]
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[third]+nums[second] == target {
				anx = append(anx, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return anx
}

//16.最接近的三数之和
func ThreeSumClosest(nums []int, target int) int {
	return 0
}

//46.全排列
func Permute(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	if n == 0 {
		return res
	}

	path := make([]int, 0)
	used := make([]bool, n)
	dfs := func(nums []int, len, depth int, path []int, used []bool, res [][]int) [][]int { return nil }
	dfs = func(nums []int, len1, depth int, path []int, used []bool, res [][]int) [][]int {
		if len1 == depth {
			res = append(res, path)
			return res
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			res = dfs(nums, len1, depth+1, path, used, res)
			path = path[:len(path)-1]
			used[i] = false
		}
		return res
	}
	res = dfs(nums, n, 0, path, used, res)
	return res
}

//53.最大子数组和
func MaxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

//54.螺旋矩阵
func SpiralOrder(matrix [][]int) []int {
	return []int{}
}
