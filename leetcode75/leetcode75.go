package leetcode75

import (
	"fmt"
	"math"
)

// 一维数组的动态和
func RunningSum(nums []int) []int {
	num := 0
	for i := 0; i < len(nums); i++ {
		aa := nums[i]
		nums[i] = nums[i] + num
		num += aa
	}
	mid := len(nums) / 2
	fmt.Println(nums[:mid-1])
	fmt.Println(nums[mid:], mid)
	return nums
}

// 寻找数组的中心下标
func PivotIndex(nums []int) int {
	index := 0
	for index < len(nums) {
		left := 0
		right := 0
		if index == 0 {
			right = array(nums[1:])
		} else if index == len(nums)-1 {
			left = array(nums[:index])
		} else {
			left = array(nums[:index])
			right = array(nums[index+1:])
		}
		if left == right {
			return index
		}
		index++
	}
	return -1
}

// 计算数组的和
func array(list []int) int {
	count := 0
	for i := 0; i < len(list); i++ {
		count += list[i]
	}
	return count
}

// 同构字符串
func IsIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}

// 判断子序列
func IsSubsequence(s string, t string) bool {
	left := 0
	right := 0
	var s1 string
	for left < len(s) && right < len(t) {
		if s[left] == t[right] {
			s1 += string(s[left])
			left++
			right++
			continue
		} else if s[left] != t[right] {
			right++
		}
	}
	fmt.Println(s1)
	return s == s1
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个有序链表
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists(list2.Next, list1)
		return list2
	}
}

// 链表的中间节点
func MiddleNode(head *ListNode) *ListNode {
	count := 0
	node := head
	for node != nil {
		node = node.Next
		count++
	}
	if count%2 == 0 {
		count = (count / 2)
	} else {
		count = count / 2
	}
	fmt.Println(count)
	node = head
	for count > 0 {
		node = node.Next
		count--
	}
	return node
}

// 环形链表二
func DetectCycle(head *ListNode) *ListNode {
	dummyhead := head
	list := make(map[*ListNode]int)
	for dummyhead != nil {
		if list[dummyhead] == 1 {
			return dummyhead
		}
		list[dummyhead] += 1
		dummyhead = dummyhead.Next
	}
	return nil
}

// 买卖股票的最佳时机
func MaxProfit(prices []int) int {
	//最大利润
	maxprofit := 0
	min := math.MaxInt32

	for i := 0; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if prices[i]-min > maxprofit {
			maxprofit = prices[i] - min
		}
	}
	return maxprofit
}

// 最长回文串
func LongestPalindrome(s string) int {
	counter := make(map[byte]int)
	count := 0
	odd := 0
	for k := range s {
		counter[s[k]]++
	}
	for _, v := range counter {
		if v%2 == 1 {
			ans := v / 2 * 2
			count += ans
			odd = 1
		} else {
			count += v
		}
	}
	return count + odd
}

// N叉树的前序遍历
type Node struct {
	Val      int
	Children []*Node
}

func Preorder(root *Node) []int {
	var dfs func(*Node)
	ans := make([]int, 0)
	dfs = func(n *Node) {
		if n == nil {
			return
		}
		ans = append(ans, n.Val)
		for _, v := range n.Children {
			dfs(v)
		}
	}
	dfs(root)
	return ans
}

// 二叉树的层序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; 0 < len(q); i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			fmt.Println(ret)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

// 二分查找
func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// 验证二叉搜索树
func IsValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}
func helper(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}
	return helper(node.Left, min, node.Val) && helper(node.Right, node.Val, max)
}

var Head *TreeNode = &TreeNode{6, left1, right1}
var left1 *TreeNode = &TreeNode{2, left2, left3}
var left2 *TreeNode = &TreeNode{0, nil, nil}
var left3 *TreeNode = &TreeNode{4, Left4, left5}
var Left4 *TreeNode = &TreeNode{3, nil, nil}
var left5 *TreeNode = &TreeNode{5, nil, nil}
var right1 *TreeNode = &TreeNode{8, right2, right3}
var right2 *TreeNode = &TreeNode{7, nil, nil}
var right3 *TreeNode = &TreeNode{9, nil, nil}

// 二叉搜索树的最近公共祖先
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	dfs := func(head, target *TreeNode) (path []*TreeNode) {
		node := head
		for node != target {
			path = append(path, node)
			if node.Val < target.Val {
				node = node.Right
			} else {
				node = node.Left
			}
		}
		path = append(path, node)
		return path
	}
	path_p := dfs(root, p)
	path_q := dfs(root, q)
	var ancestor *TreeNode
	for i := 0; i < len(path_p) && i < len(path_q) && path_p[i] == path_q[i]; i++ {
		ancestor = path_p[i]
	}
	return ancestor
}

// 图像渲染
func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	R, C := len(image), len(image[0])
	oldColor := image[sr][sc]
	if oldColor == color {
		return image
	}
	dfs := func(r, c int) {}
	dfs = func(r, c int) {
		if image[r][c] == oldColor {
			image[r][c] = color
			if r >= 1 {
				dfs(r-1, c)
			}
			if r+1 < R {
				dfs(r+1, c)
			}
			if c >= 1 {
				dfs(r, c-1)
			}
			if c+1 < C {
				dfs(r, c+1)
			}
		}
	}
	dfs(sr, sc)
	return image
}

// 岛屿的数量
func NumIslands(grid [][]byte) int {
	nr := len(grid)
	nc := len(grid[0])
	dfs := func(r, c int) {}
	dfs = func(r, c int) {
		grid[r][c] = 48
		if r-1 >= 0 && grid[r-1][c] == 49 {
			dfs(r-1, c)
		}
		if r+1 < nr && grid[r+1][c] == 49 {
			dfs(r+1, c)
		}
		if c-1 >= 0 && grid[r][c-1] == 49 {
			dfs(r, c-1)
		}
		if c+1 < nc && grid[r][c+1] == 49 {
			dfs(r, c+1)
		}
	}
	count := 0
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == 49 {
				count += 1
				dfs(i, j)
			}
		}
	}
	return count
}

// 斐波那契数
func Fib(n int) int {
	nums := make(map[int]int)
	nums[0] = 0
	nums[1] = 1
	for i := 2; i <= 30; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums[n]
}

// 爬楼梯
func ClimbStairs(n int) int {
	memo := make([]int, n+1)
	dfs := func(n int) int { return 0 }
	dfs = func(n int) int {
		if memo[n] > 0 {
			return memo[n]
		}
		if n == 1 {
			memo[n] = 1
		} else if n == 2 {
			memo[n] = 2
		} else {
			memo[n] = dfs(n-1) + dfs(n-2)
		}
		return memo[n]
	}
	return dfs(n)
}

// 找到字符串中所有字母异位词
func FindAnagrams(s string, p string) []int {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return nil
	}
	ans := []int{}
	var sCount, pCount [26]int
	for k, v := range p {
		sCount[s[k]-'a'] += 1
		pCount[v-'a'] += 1
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}
	for k, v := range s[:sLen-pLen] {
		sCount[v-'a'] -= 1
		sCount[s[k+pLen]-'a'] += 1
		if sCount == pCount {
			ans = append(ans, k+1)
		}
	}
	return ans
}

// 替换后的最长重复字符串
func CharacterReplacement(s string, k int) int {
	cnt := [26]int{}
	maxCnt, left := 0, 0
	maxDfs := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for right, ch := range s {
		cnt[ch-'A']++
		maxCnt = maxDfs(maxCnt, cnt[ch-'A'])
		if right-left+1-maxCnt > k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

// 比较含退格的字符串
func BackspaceCompare(s string, t string) bool {
	index1 := 0
	index2 := 0
	sLen, tLen := len(s), len(t)
	s1, s2 := "", ""
	for index1 < sLen || index2 < tLen {
		if index1 < sLen {
			if string(s[index1]) != "#" {
				s1 += string(s[index1])
			} else {
				if len(s1) != 0 {
					s1 = s1[:len(s1)-1]
				}
			}
			index1++
		}
		if index2 < tLen {
			if string(t[index2]) != "#" {
				s2 += string(t[index2])
			} else {
				if len(s2) != 0 {
					s2 = s2[:len(s2)-1]
				}
			}
			index2++
		}
	}
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	return s1 == s2
}

// 快乐数
func IsHappy(n int) bool {
	return false
}
