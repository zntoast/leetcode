package leetcodehot100

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两数之和
func TwoSum(nums []int, target int) []int {
	length := len(nums)
	index := make([]int, 2)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				index[0] = i
				index[1] = j
			}
		}
	}
	return index
}

// 两数相加
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var left *ListNode
	var right *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum/10, sum%10
		if left == nil {
			left = &ListNode{sum, nil}
			right = left
		} else {
			right.Next = &ListNode{sum, nil}
			right = right.Next
		}
	}
	if carry >= 0 {
		right.Next = &ListNode{carry, nil}
	}
	return left
}

// 无重复字符的最长子串
func LengthOfLongestSubstring(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		list := make(map[byte]int)
		list[s[i]] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] != s[j] {
				if list[s[j]] == 1 {
					break
				}
				list[s[j]] = list[s[j]] + 1
			} else {
				break
			}
		}
		if len(list) > count {
			count = len(list)
		}
	}
	return count
}

// 寻找两个正序数组的中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	list := make([]int, 0)
	list = append(list, nums1...)
	list = append(list, nums2...)
	sort.Ints(list)
	mid := len(list) / 2
	if len(list)%2 == 1 {
		return float64(list[mid])
	} else {
		return (float64(list[mid]) + float64(list[mid-1])) / 2
	}
}

// 最长回文子串
func LongestPalindrome(s string) string {
	return ""
}

// 最长公共子序列
func LongestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, len1+1)
	for k := range dp {
		dp[k] = make([]int, len2+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[len1][len2]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 回文数
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
	fmt.Println(newNum)
	fmt.Println(aa)
	fmt.Println(newNum == aa)
	return newNum == aa

}

// 10.正则表达式匹配
func IsMatch(s string, p string) bool {
	reg, err := regexp.Compile(p)
	if err != nil {
		log.Fatal(err)
	}
	new := reg.FindString(s)
	fmt.Println(new)
	return new == s
}

// 11.盛最多水的容器
func MaxArea(height []int) int {
	l := 0
	r := len(height) - 1
	anx := 0
	for l < r {
		arer := min(height[l], height[r]) * (r - l)
		anx = max(arer, anx)
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return anx
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 15.三数之和
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

// 最接近的三数之和
func ThreeSumClosest(nums []int, target int) int {
	return 0
}

var list []string
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// 17.电话号码和字母的组合
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	list = []string{}
	backtrack(digits, 0, "")
	return list
}
func backtrack(digits string, index int, s string) {
	if index == len(digits) {
		list = append(list, s)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		length := len(letters)
		for i := 0; i < length; i++ {
			backtrack(digits, index+1, s+string(letters[i]))
		}
	}
}

// 删除链表的倒数第N个节点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	rightHead := head
	leftHead := head
	for i := 0; i < n; i++ {
		rightHead = rightHead.Next
	}
	if rightHead == nil {
		return head.Next
	}
	for rightHead.Next != nil {
		leftHead = leftHead.Next
		rightHead = rightHead.Next
	}
	leftHead.Next = leftHead.Next.Next
	return head
}

// 33. 搜索旋转排序数组
func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid + 1
			} else {
				left = mid - 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 移除元素
func RemoveElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	nums = nums[:left]
	fmt.Println(nums)
	return left
}

// 串联所有单词的子串
func FindSubstring(s string, words []string) []int {
	ans := []int{}
	flag := make([]bool, len(words))
	m := make(map[int]int)
	dfs := func(s, str string, index int, words []string) {}
	dfs = func(s, str string, index int, words []string) {
		if index == len(words) {
			num := strings.Index(s, str)
			if num != -1 && m[num] == 0 {
				m[num]++
				ans = append(ans, num)
			}
			return
		}
		for i := 0; i < len(words); i++ {
			if !flag[i] {
				flag[i] = true
				str += words[i]
				dfs(s, str, index+1, words)
				flag[i] = false
				str = str[:(len(str) - len(words[i]))]
			}
		}
	}
	dfs(s, "", 0, words)
	return ans
}
