package medium

import (
	"bytes"
	"container/heap"
	"container/list"
	"fmt"
	"math"
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
		// 记录字母最后一次出现位置的下标
		temp[s[i]] = i
	}
	res := make([]int, 0)
	left := 0
	right := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for j := 0; j < len(s); j++ {
		// 不到刷新字母最后出现的位置的下标
		right = max(right, temp[s[j]])
		if right == j {
			res = append(res, right-left+1)
			left = right + 1
		}
	}
	return res
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
	// 罗马数字对应的阿拉伯数字
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
	// 暂时不写
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

/*
LC 55. 跳跃游戏
*/
func CanJump(nums []int) bool {
	n := len(nums)
	last := n - 1
	for i := n - 2; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}
	return last == 0
}

/*
LC 	77. 组合
*/
func Combine(n int, k int) [][]int {
	ans := [][]int{}
	temp := []int{}
	var dfs func(cur int)
	dfs = func(cur int) {
		if len(temp)+(n-cur+1) < k {
			return
		}
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return ans
}

/*
lc 	45. 跳跃游戏 II
*/
func Jump(nums []int) int {
	position := len(nums) - 1
	res := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				res++
				position = i
				break
			}
		}
	}
	return res
}

/*
LC 43. 字符串相乘
*/
func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := ""
	m, n := len(num1), len(num2)
	ansArr := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i] - '0')
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ansArr[i+j+1] += x * y
		}
	}
	for i := m + n - 1; i > 0; i-- {
		ansArr[i-1] += ansArr[i] / 10
		ansArr[i] %= 10
	}
	index := 0
	if ansArr[0] == 0 {
		index = 1
	}
	for ; index < m+n; index++ {
		ans += strconv.Itoa(ansArr[index])
	}
	return ans
}

/*
LC 1488. 避免洪水泛滥
*/
func AvoidFlood(rains []int) []int {
	return nil
}

/*
LC 260. 只出现一次的数字 III
*/
func SingleNumber(nums []int) []int {
	exist := make(map[int]bool)
	for _, v := range nums {
		if _, ok := exist[v]; ok {
			exist[v] = false
		} else {
			exist[v] = true
		}
	}
	ans := make([]int, 0)
	for k, v := range exist {
		if v {
			ans = append(ans, k)
		}
	}
	return ans
}

/*
LC 39. 组合总和
*/
func CombinationSum(candidates []int, target int) [][]int {
	temp := []int{}
	ans := [][]int{}
	dfs := func(target, idx int) {}
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int{}, temp...))
			return
		}
		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			temp = append(temp, candidates[idx])
			dfs(target-candidates[idx], idx)
			temp = temp[:len(temp)-1]
		}
	}
	dfs(target, 0)
	return ans
}

/*
LC 2400. 恰好移动 k 步到达某一位置的方法数目
*/
func NumberOfWays(startPos int, endPos int, k int) int {
	mod := 1000000007
	d := abs(startPos - endPos)
	if (d+k)%2 == 1 || d > k {
		return 0
	}
	f := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		f[i] = make([]int, k+1)
		f[i][0] = 1
		for j := 1; j <= i; j++ {
			f[i][j] = (f[i-1][j] + f[i-1][j-1]) % mod
		}
	}
	return f[k][(d+k)/2]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LC 1261. 在受污染的二叉树中查找元素
type FindElements struct {
	Val map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	element := FindElements{
		Val: map[int]bool{
			0: true,
		},
	}
	dfs := func(node *TreeNode, val int64) {}
	dfs = func(node *TreeNode, val int64) {
		if node == nil {
			return
		}
		if node.Left != nil {
			temp := val*2 + 1
			element.Val[int(temp)] = true
			dfs(node.Left, temp)
		}
		if node.Right != nil {
			temp := val*2 + 2
			element.Val[int(temp)] = true
			dfs(node.Right, temp)
		}
	}
	dfs(root, 0)
	return element
}

func (this *FindElements) Find(target int) bool {
	return this.Val[int(target)]
}

// LC 2789. 合并后数组中的最大元素
func MaxArrayValue(nums []int) int64 {
	var dfs func(nums []int)
	dfs = func(nums []int) {
		if len(nums) < 2 {
			return
		}
		for i := len(nums) - 1; i >= 0; i-- {
			if i-1 >= 0 && nums[i] >= nums[i-1] {
				nums[i] += nums[i-1]
				dfs(append(nums[:i-1], nums[i:]...))
				return
			}
		}
	}
	dfs(nums)
	return int64(nums[0])
}

// LC 2831. 找出最长等值子数组
func LongestEqualSubarray(nums []int, k int) int {
	ans := 0
	pos := make(map[int][]int)
	for k, num := range nums {
		pos[num] = append(pos[num], k)
	}
	mas := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for _, p := range pos {
		j := 0
		for i := 0; i < len(p); i++ {
			for p[i]-p[j]-(i-j) > k {
				j++
			}
			ans = mas(ans, i-j+1)
		}
	}
	return ans
}

// 2981. 找出出现至少三次的最长特殊子字符串 I
func MaximumLength(s string) int {
	tempMap := map[string]int{}
	for k := range s {
		s1 := string(s[k])
		tempMap[s1]++
		if k+1 < len(s) && s[k] == s[k+1] {
			s2 := string(s[k+1])
			tempMap[s1+s2]++
		}
		if k+2 < len(s) && s[k] == s[k+1] && s[k+1] == s[k+2] {
			s2 := string(s[k+1])
			s3 := string(s[k+2])
			tempMap[s1+s2+s3]++
		}
	}
	type ans struct {
		key string // 最长特殊子字符串
		val int    // 出现次数
	}
	a := ans{"", -1}
	for ss, v := range tempMap {
		if v < 3 {
			continue
		}
		if len(ss) < len(a.key) {
			continue
		}
		if len(ss) == len(a.key) {
			if a.val < v {
				a.key = ss
				a.val = v
			}
		} else {
			a.key = ss
			a.val = v
		}
	}
	return a.val
}

// LC 2957. 消除相邻近似相等字符
func RemoveAlmostEqualCharacters(word string) int {
	ans := 0
	i := 1
	for i < len(word) {
		left := word[i-1]
		right := word[i]
		temp := 0
		if left > right {
			temp = int(left) - int(right)
		} else {
			temp = int(right) - int(left)
		}
		if temp > 1 {
			i++
		} else {
			ans++
			i += 2
		}
	}
	return ans
}

// LC 2101. 引爆最多的炸弹
func MaximumDetonation(bombs [][]int) int {
	abs := func(a int) int {
		if a > 0 {
			return a
		} else {
			return -a
		}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	bombMap := map[int][]int{}
	// 判断每个点是否可以引爆其他节点
	for i := 0; i < len(bombs); i++ {
		for j := 0; j < len(bombs); j++ {
			// 节点重合
			if i == j {
				continue
			}
			// 两节点的距离
			var diff float64 = 0
			length := abs(bombs[i][0] - bombs[j][0])
			width := abs(bombs[i][1] - bombs[j][1])
			if width == 0 { // 同行
				diff = float64(length)
			} else if length == 0 { // 同列
				diff = float64(width)
			} else { // 不同行不同列
				diff = math.Sqrt(float64(length*length + width*width))
			}
			if diff <= float64(bombs[i][2]) {
				bombMap[i] = append(bombMap[i], j)
			}
		}
	}
	dfs := func(nums []int, hasBombs map[int]bool, temp int) int { return 0 }
	dfs = func(nums []int, hasBombs map[int]bool, temp int) int {
		for _, num := range nums {
			if hasBombs[num] {
				continue
			}
			hasBombs[num] = true
			temp++
			temp += dfs(bombMap[num], hasBombs, 0)
		}
		return temp
	}
	ans := 1
	for k, nums := range bombMap {
		ans = max(ans, dfs(nums, map[int]bool{k: true}, 1))
	}
	return ans
}

// LC 2844. 生成特殊数字的最少操作
func MinimumOperations(num string) int {
	f0, f5 := false, false
	n := len(num)
	for i := n - 1; i >= 0; i-- {
		if num[i] == '0' || num[i] == '5' {
			if f0 {
				return n - i - 2
			}
			if num[i] == '0' {
				f0 = true
			}
			if num[i] == '5' {
				f5 = true
			}
		} else if num[i] == '2' || num[i] == '7' {
			if f5 {
				return n - i - 2
			}
		}
	}
	if f0 {
		return n - 1
	}
	return n
}

// LC 3128. 直角三角形
func NumberOfRightTriangles(grid [][]int) int64 {
	var ans int64 = 0
	// 计算每行每列 中 1的数量
	rowMap, colMap := map[int]int{}, map[int]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				rowMap[i]++
				colMap[j]++
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				ans += int64(rowMap[i]-1) * int64(colMap[j]-1)
			}
		}
	}
	return ans
}

// LC 3129. 找出所有稳定的二进制数组 I
func NumberOfStableArrays(zero int, one int, limit int) int {
	ans := 0
	l := zero + one
	dfs := func(zero int, one int, v int, last int, num int, arr []int) {}
	dfs = func(zero, one int, v int, last int, num int, arr []int) {
		if last == v {
			num++
		} else {
			num = 0
		}
		arr = append(arr, v)
		if num == limit {
			return
		}
		if len(arr) == l {
			ans++
			return
		}
		if zero > 0 {
			dfs(zero-1, one, 0, v, num, arr)
		}
		if one > 0 {
			dfs(zero, one-1, 1, v, num, arr)
		}
	}
	dfs(zero-1, one, 0, 1, 0, make([]int, 0, l))
	dfs(zero, one-1, 1, 0, 0, make([]int, 0, l))
	return ans
}

// LC 676. 实现一个魔法字典
type MagicDictionary struct {
	wordLen map[int][]string
}

func Constructor1() MagicDictionary {
	return MagicDictionary{
		wordLen: make(map[int][]string),
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, v := range dictionary {
		this.wordLen[len(v)] = append(this.wordLen[len(v)], v)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for _, word := range this.wordLen[len(searchWord)] {
		left := 0
		right := len(word) - 1
		temp := 0
		for left <= right {
			if word[left] != searchWord[left] && left != right {
				temp++
			}
			if word[right] != searchWord[right] && left != right {
				temp++
			}
			if word[right] != searchWord[right] && left == right {
				temp++
			}
			if temp > 1 {
				break
			}
			left++
			right--
		}
		if temp == 1 {
			return true
		}
	}
	return false
}

// LC 3152. 特殊数组 II
func IsArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		if (nums[i]^nums[i-1])&1 == 1 {
			dp[i] = dp[i-1] + 1
		}
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		x, y := q[0], q[1]
		res[i] = dp[y] >= y-x+1
	}
	return res
}

// LC 3143. 正方形中的最多点数
func MaxPointsInsideSquare(points [][]int, s string) int {
	ans := 0
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	// 同一个标签在坐标轴上的分布情况
	stemp := map[byte][]int{}
	for k, point := range points {
		length := 0
		if abs(point[0]) > abs(point[1]) {
			length = abs(point[0])
		} else {
			length = abs(point[1])
		}
		stemp[s[k]] = append(stemp[s[k]], length)
	}
	max := math.MaxUint32
	// 计算正方形的最大长度
	for _, v := range stemp {
		if len(v) < 2 {
			continue
		}
		sort.Ints(v)
		var l int
		if v[0] >= v[1] {
			l = v[0] - 1
		} else {
			l = v[1] - 1
		}

		if max > l {
			max = l
		}
	}
	// 统计点在正方形中的数据
	for _, nums := range stemp {
		for _, v := range nums {
			if v <= max {
				ans++
			}
		}
	}
	return ans
}

// LC 2554. 从一个范围内选择最多整数 I
func MaxCount(banned []int, n int, maxSum int) int {
	sort.Ints(banned)
	exist := map[int]bool{}
	for i := range banned {
		exist[banned[i]] = true
	}
	pass := []int{}
	for i := 1; i <= n; i++ {
		if exist[i] {
			continue
		}
		pass = append(pass, i)
	}
	ans := 0
	num := 0
	for _, v := range pass {
		if num+v > maxSum {
			break
		}
		ans++
		num += v
	}
	return ans
}

// LC 1845. 座位预约管理系统
type SeatManager struct {
	available *Heap
}

func NewConstructor(n int) SeatManager {
	h := &Heap{}
	heap.Init(h)
	for i := 1; i <= n; i++ {
		heap.Push(h, i)
	}
	return SeatManager{available: h}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.available).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.available, seatNumber)
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// LC 416. 分割等和子集
func CanPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	max := 0
	for i := range nums {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}
	target := sum / 2
	if sum%2 == 1 {
		return false
	}
	if max > target {
		return false
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}

	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n-1][target]
}
