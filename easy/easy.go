package easy

import (
	"sort"
	"strconv"
)

/*
1005. K 次取反后最大化的数组和
*/
func LargestSumAfterKNegations(nums []int, k int) int {
	min_index := 0
	result := 0
	sort.Ints(nums)
	for i := range nums {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
		if nums[min_index] > nums[i] {
			min_index = i
		}
	}
	if k%2 == 1 {
		nums[min_index] *= -1
	}
	for _, v := range nums {
		result += v
	}
	return result
}

/*
LC 13. 罗马数字转整数
*/
func RomanToInt(s string) int {
	//罗马数字对应的阿拉伯数字
	symbolValues := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	ans := 0
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		//下一个字符对应的数值大于当前字符对应数值
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}

/*
LC 35.搜索插入位置
*/
func SearchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < target {
			continue
		}
		return i
	}
	return len(nums)
}

/*
LC 58.最后一个单词的长度
*/
func LengthOfLastWord(s string) int {
	lenght := len(s)
	ans := 0
	for i := lenght - 1; i >= 0; i-- {
		if s[i] == ' ' && ans == 0 {
			continue
		}
		if s[i] == ' ' {
			return ans
		}
		ans++
	}
	return ans
}

/*
LC 67.二进制求和
*/
func AddBinary(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := max(lenA, lenB)
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

/*
LC 69.x的平方根
*/
func MySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l < r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 链
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
LC 83.删除排序链表中的重复元素
*/
func DeleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
LC 100.相同的树
*/
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	ans := true
	// 深度优先遍历
	dfs := func(p *TreeNode, q *TreeNode) {}
	dfs = func(p, q *TreeNode) {
		if p == nil && q == nil {
			return
		} else if p != nil && q != nil {
			if p.Val != q.Val {
				ans = false
			}
			dfs(p.Left, q.Left)
			dfs(p.Right, q.Right)
		} else {
			ans = false
		}
	}
	dfs(p, q)
	return ans
}

/*
LC 88.合并两个有序数组
*/
func Merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			i--
			m--
			nums1[i] = nums1[m]
		} else {
			i--
			n--
			nums1[i] = nums2[n]
		}
	}
}

/*
LC 101. 对称二叉树
*/
func IsSymmetric(root *TreeNode) bool {
	isMirror := func(l, r *TreeNode) bool { return false }
	isMirror = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		return l.Val == r.Val && isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
	}
	return isMirror(root, root)
}

/*
LC 104. 二叉树的最大深度
深度优先算法
*/
func MaxDepth(root *TreeNode) int {
	ans := 0
	dfs := func(r *TreeNode, depth int) {}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dfs = func(r *TreeNode, depth int) {
		if r == nil {
			return
		}
		ans = max(depth, ans)
		dfs(r.Left, depth+1)
		dfs(r.Right, depth+1)
	}
	dfs(root, 1)
	return ans
}

/*
LC 108.将有序数组转换为二叉搜索树
*/
func SortedArrayToBST(nums []int) *TreeNode {
	dfs := func(nums []int, l, r int) *TreeNode { return nil }
	dfs = func(nums []int, l, r int) *TreeNode {
		if l < r {
			return nil
		}
		mid := l + (r-l)/2
		node := &TreeNode{nums[mid], nil, nil}
		node.Left = dfs(nums, l, mid-1)
		node.Right = dfs(nums, mid+1, r)
		return node
	}
	return dfs(nums, 0, len(nums)-1)
}
