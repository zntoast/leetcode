package intermediate

import (
	"math"
	"sort"
	"strings"
)

/*
LC 三数之和
*/
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for first := 0; first < n; first++ {
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
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

/*
LC 矩阵置零
*/
func SetZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	row := make([]bool, m)
	col := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}

/*
LC 字母异位词分组
*/
func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	m := make([][]string, 0)
	temp := map[string]int{}
	index := 0
	for _, s := range strs {
		w := strings.Split(s, "")
		sort.Strings(w)
		s1 := strings.Join(w, "")
		if v, ok := temp[s1]; ok {
			m[v] = append(m[v], s)
		} else {
			m = append(m, []string{})
			m[index] = append(m[index], s)
			temp[s1] = index
			index++
		}
	}
	return m
}

/*
LC 无重复字符的最长子串
*/
func LengthOfLongestSubstring(s string) int {
	result := 0
	for k := range s {
		temp := map[byte]bool{}
		l := 1
		temp[s[k]] = true
		for i := k + 1; i < len(s); i++ {
			if temp[s[i]] {
				break
			}
			temp[s[i]] = true
			l++
		}
		if len(temp) > result {
			result = len(temp)
		}
	}
	return result
}

/*
LC 最长回文子串
*/
func LongestPalindrome(s string) string {
	start, end := 0, 0
	context := func(left, right int) (int, int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return left + 1, right - 1
	}
	for k := range s {
		left1, right1 := context(k, k)
		left2, right2 := context(k, k+1)
		if right1-left1 > end-start {
			start = left1
			end = right1
		}
		if right2-left2 > end-start {
			start = left2
			end = right2
		}
	}
	return s[start : end+1]
}

/*
LC 递增的三元子序列
*/
func IncreasingTriplet(nums []int) bool {
	min := math.MaxInt32
	mid := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if nums[i] <= min {
			min = nums[i]
		} else if nums[i] <= mid {
			mid = nums[i]
		} else {
			return true
		}
	}
	return false
}
