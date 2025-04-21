package interview

import (
	"container/heap"
	"fmt"
	"leetcode/array"
	"math"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"
)

// 给定一个数组，数组全是整数。要求将奇数全部按顺序排在数组前面，偶数排在后面，不能新建其它数组
// 题解一、
func Array1(list []int) []int {
	length := 0
	for i := 0; i < len(list); i++ {
		if list[i]%2 == 1 {
			length++
		}
	}
	left := 0
	right := length
	for left < length && right < len(list) {
		if list[left]%2 == 0 && list[right]%2 == 1 {
			aa := list[left]
			list[left] = list[right]
			list[right] = aa
			left++
			right++
			continue
		}
		if list[left]%2 == 1 {
			left++
		}
		if list[right]%2 == 0 {
			right++
		}
	}
	array.Sort(list[:length])
	array.Sort(list[length:])
	return list
}

// 题解二、
func Array(list []int) []int {
	for i := 0; i < len(list); i++ {
	}
	return nil
}

// 给定一个数组1表示种有花,0表示没有。花与花之间不能紧挨着种。n表示可以种植的数量，可以则返回true，否则false
func Flower(list []int, n int) bool {
	left := 0
	right := 0
	m := 0
	for i := 0; i < len(list); i++ {
		if list[i] == 0 {
			left = i - 1
			right = i + 1
			if left < 0 {
				if list[right] == 0 {
					m++
					i++
				}
			} else if left >= 0 && right < len(list) {
				if list[left] == 0 && list[right] == 0 {
					m++
					i++
				}
			} else if right >= len(list) {
				if list[left] == 0 {
					m++
					i++
				}
			}
		}
	}
	fmt.Println(m)
	return m >= n
}

// 找出两个字符串的最大公共交集
func FindString(s1, s2 string) string {
	str := ""
	if len(s1) < len(s2) {
		str = do(s1, s2)
	} else {
		str = do(s2, s1)
	}
	return str
}

func do(min, max string) string {
	len1 := len(min)
	str := ""
	for i := 0; i < len1; i++ {
		s := string(min[i])
		index := i
		for strings.Contains(max, s) {
			if len(str) < len(s) {
				str = s
			}
			if index < len1-1 {
				s = s + string(min[index+1])
			} else {
				break
			}
			index++
		}
	}
	return str
}

/*
斐波那契数列
*/
func FibonacciSequence(num int) {
	start := time.Now()
	nums := make([]*big.Int, num)
	for i := 0; i < num; i++ {
		if i <= 1 {
			nums[i] = big.NewInt(1)
		} else {
			temp := new(big.Int)
			nums[i] = temp.Add(nums[i-1], nums[i-2])
		}
		fmt.Printf("数位第%v位: %v\n", i, nums[i])
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("执行所耗: %v\n", delta)
}

/*
快速排序
*/
func FastSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// mid中间值(假定mid = values[0] 为中间值)	, i下标
	mid, i := nums[0], 1
	left, right := 0, len(nums)-1
	for left < right {
		// 把 小于mid 的值放到 head前面 反则亦然
		if nums[i] > mid {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		}
	}
	// 中间值替换成mid
	nums[left] = mid
	FastSort(nums[:left])
	FastSort(nums[left+1:])
}

/*
冒泡排序
*/
func BubbleSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	time.Sleep(time.Nanosecond * 100)
}

/*
插入排序
*/
func InsertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1], nums[j] = nums[j], key
			j--
		}
	}
}

/*
合并排序 将两个或多个有序数组进行排序
*/
func MergeSort(a []int, left, right int) {
	if left == right {
		return
	}
	if left < right {
		// 拆分成n份进行排序
		mid := (left + right) / 2
		MergeSort(a, left, mid)
		MergeSort(a, mid+1, right)
		Merge(a, left, right)
	}
}

// 合并数组
func Merge(a []int, left, right int) {
	m := right - left + 1
	// 临时数组b
	b := make([]int, m)
	left0 := left
	// 数组下标
	i := 0
	// 中间数组的中间值
	mid := (left + right) / 2
	k := mid + 1
	for left <= mid && k <= right {
		if a[left] < a[k] {
			b[i] = a[left]
			i++
			left++
		} else {
			b[i] = a[k]
			i++
			k++
		}
	}
	if left > mid {
		for k <= right {
			b[i] = a[k]
			i++
			k++
		}
	}
	if k > right {
		for left <= mid {
			b[i] = a[left]
			left++
			i++
		}
	}
	for j := 0; j < m; j++ {
		a[left0] = b[j]
		left0++
	}
}

/*
统计n位数以内有多少个素数,不计入 0 ,1 ,埃筛法
*/
func Eratosthenes(n int) int {
	if n < 2 {
		fmt.Println("n 不能小于2")
		return n
	}
	count := 0
	// 素数存储
	isPrims := make([]bool, n)
	for i := 2; i < n; i++ {
		// 判断是否为素数 ,是则进入
		if !isPrims[i] {
			count++
			for j := 2 * i; j < n; j += i {
				// 将非素数改为true
				isPrims[j] = true
			}
		}
	}
	return count
}

/*
删除有序数组的重复项,每个元素只能出现一次,返回删除后数组的长度
考察双指针
*/
func RemoveDuplicates(nums []int) int {
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

/*
寻找数组的中心下标,使得左边数组的和跟右边数组的和相等
*/
func PivotIndex(nums []int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		if total == num {
			return i
		}
		num -= nums[i]
	}
	return -1
}

/*
在不使用官方包情况, 算x的平方根 取整数
二分查找
*/
func BinarySearch(x int) int {
	if x < 2 {
		return 1
	}
	index, left, right := -1, 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			index = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return index
}

type ListNote struct {
	Val  int
	Next *ListNote
}

/*
反转链表,递归实现
*/
func Recursion(head *ListNote) *ListNote {
	if head == nil || head.Next == nil {
		return head
	}
	node := Recursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

/*
数组三个数的最大乘积
*/
func Sort(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(nums[0]*nums[1]*nums[n-1], nums[n-1]*nums[n-2]*nums[n-3])
}

/*
两数之和，无序数组,获取两个数的下标(用数组返回)
*/
func Solution(nums []int, tager int) []int {
	isExist := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		a := tager - nums[i]
		if _, ok := isExist[a]; ok {
			return []int{isExist[a], i}
		}
		isExist[nums[i]] = i
	}
	return nil
}

/*
排列硬币,摆成阶梯,第k行有k枚硬币 .n枚硬币能摆多少行(完成的)
*/
func ArrangeCoins(n int) (int, int) {
	if n < 1 {
		return 0, 0
	}
	count := 0
	for k := 1; k < n; k++ {
		count += k
		if ((k+1)*k/2 <= n) && ((k+2)*(k+1)/2 > n) {
			return k, count
		}
	}
	return -1, 0
}

/*
寻找符合要求的字符串
*/
func FindStrings(a, b string) int {
	count := 0
	l := len(b)
	for i := l; i <= len(a); i++ {
		if a[i-l:i] == b {
			count += 1
		}
	}
	return count
}

/*
华为第一题
*/
func test1() {
	n := 0
	ans := -1
	values := map[int][]int{}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scan(&x)
		values[x] = append(values[x], i)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for _, v := range values {
		sort.Ints(v)
		if len(v) >= 2 {
			ans = max(ans, v[len(v)-1]-v[0])
		}
	}
	fmt.Printf("%d\n", ans)
}

/*
华为面试第二题 (未完成)
*/
func test2() {
	n := 0
	ans := ""
	j := 0
	res := map[int]string{}
	invert := func(s string) string {
		result := ""
		for k := range s {
			result = result + string(s[k])
		}
		return result
	}
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		x := ""
		fmt.Scan(&x)
		if x == " " {
			// 反转
			ans += invert(res[j])
			j++
		}
		res[j] += x
	}
	fmt.Printf("%s\n", ans)
}

/*
华为面试第三题 Excel单元格数值统计
*/
func test3() {
	// table := map[int][]string{}
	// a := 0
	// b := 0
	// for {
	// 	n, _ := fmt.Scan(&a, &b)
	// 	if n == 0 {
	// 		break
	// 	} else {
	// 		fmt.Printf("%d\n", a+b)
	// 	}
	// }
}

/*
面试题67. 把字符串转换成整数
*/
func StrToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	ans := 0
	isSigna := 0
	sign := 1
	// 删除前面的空格
	for str[0] == ' ' && len(str) >= 2 {
		str = str[1:]
	}
	if str == " " {
		return 0
	}
	for i := 0; i < len(str); i++ {
		// 传入的字符 不是数值
		if str[i] < '0' || str[i] > '9' {
			// 是否为+号
			isAdd := str[i] == '+'
			// 是否为-号
			isMinus := str[i] == '-'
			// 不是加号也不是减号
			if !isAdd && !isMinus {
				break
			}
		}
		if str[i] == '-' || str[i] == '+' {
			isSigna++
			if str[i] == '-' {
				sign = -1
			}
			if isSigna > 1 {
				return 0
			}
			continue
		}
		ans = ans*10 + int((str[i] - '0'))
	}
	ans = ans * sign
	if ans < math.MinInt32 {
		return math.MinInt32
	} else if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	return ans
}

/*
遍历二叉树
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func QueryTree() []int {
	head := &TreeNode{}
	ans := []int{}
	dfs := func(node *TreeNode) {}
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(head)
	return ans
}

// 华为机试题
// 给定一个非空字符串S，其被N个‘-’分隔成N+1的子串，给定正整数K，
// 要求除第一个子串外，其余的子串每K个字符组成新的子串，并用‘-’分隔。
// 对于新组成的每一个子串，如果它含有的小写字母比大写字母多，
// 则将这个子串的所有大写字母转换为小写字母；反之，如果它含有的大写字母比小写字母多，
// 则将这个子串的所有小写字母转换为大写字母；大小写字母的数量相等时，不做转换
// 示例1 输入
// 3
// 12abc-abCABc-4aB@
// 输出
// 12abc-abc-ABC-4aB-@
func HuaWeiStringHanle(S string, K int) string {
	s := strings.Split(S, "-")
	news := ""
	if len(s) == 0 {
		return ""
	}
	if len(s) > 1 {
		news = strings.Join(s[1:], "")
	}
	temp := ""
	isUpNum := 0
	isLowNum := 0
	result := s[0]
	for k, v := range news {
		temp += string(v)
		if 'a' <= v && v <= 'z' {
			isLowNum++
		}
		if 'A' <= v && v <= 'Z' {
			isUpNum++
		}
		if len(temp) == K || k == len(news)-1 {
			if isUpNum > isLowNum {
				temp = strings.ToUpper(temp)
			} else if isUpNum < isLowNum {
				temp = strings.ToLower(temp)
			}
			result += "-" + temp
			temp = ""
			isUpNum = 0
			isLowNum = 0
		}
	}
	return result
}

// 给定一个整数n，你可以进行以下三种操作
// 操作1: +1
// 操作2;+2
// 操作3: x2
// 问最少需要多少次操作可以将 0 转为为 n。
func MinOpt(n int) int {
	ans := 0
	for n >= 2 {
		n = n / 2
		ans++
	}
	ans++
	return ans
}

// 华为外包机考
// 题目描述
// 某个打印机根据打印队列执行打印任务。打印任务分为九个优先级，分别用数字1-9表示，数字越大优先级越高。打印机每次从队列头部取出第一个任务A，
// 然后检查队列余下任务中有没有比A优先级更高的任务，如果有比A优先级高的任务，则将任务A放到队列尾部，否则就执行任务A的打印。
// 请编写一个程序，根据输入的打印队列，输出实际的打印顺序
func HuaWeiQueuePrint(tasks []int) []int {
	length := len(tasks)
	h := &Heap{}
	for i := 0; i < length; i++ {
		heap.Push(h, []int{i, tasks[i]})
	}
	heapLen := h.Len()
	order := 0
	result := make([]int, length)
	for i := 0; i < heapLen; i++ {
		task := heap.Pop(h).([]int)
		index := task[0]
		result[index] = order
		order++
	}
	return result
}

func HuaWeiQueuePrint1(t []int) []int {
	tasks := make([]Task, len(t))
	for i := 0; i < len(t); i++ {
		tasks[i] = Task{i, t[i]}
	}

	result := make([]int, len(tasks))
	queue := make([]Task, len(tasks))
	copy(queue, tasks)
	order := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		hasHigher := false
		for _, t := range queue {
			if t.priority > current.priority {
				hasHigher = true
				break
			}
		}
		if hasHigher {
			queue = append(queue, current)
		} else {
			result[current.Index] = order
			order++
		}
	}
	return result
}

type Task struct {
	Index    int
	priority int
}

type Heap [][]int

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	key1, key2 := h[i][0], h[j][0]
	value1, value2 := h[i][1], h[j][1]
	if value1 < value2 {
		return false
	} else if value1 > value2 {
		return true
	} else {
		return key1 < key2
	}
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 数组排排列求和
// 我们有一个不包含重复元素且长度为3的倍数的数组，可以执行以下操作任意次：
// 每次操作可以移动数组中的任意一个元素到任意位置。
// 为使得数组和 s 达到最大可能的值，最少需要执行多少次操作？
// 注：数组和s定义如下：将数组中的元素每三个一组进行分组，对于每个三个元素的组，
// 去掉其中的最小值和最大值，保留中间值（即第二大的数）。将所有保留的中间值相加，得到总和 s。
// 例如 数组{ 1，5，3，6，2，4}，如果不移动元素，s = 3+4 =7；
// 我们可以把3移动到6和2之间得到数组{ 1，5，6，3，2，4}， 此时得到最大的s，s=5+3=8
func SummationOfArrayPermutations(nums []int) int {
	return 0
}

// 例 ： {1，2，3，4，5，6，7，8，9} 永远拿最大数+最小数+倒数第二大 依此类推 {1,8,9} {2,6,7},{3,4,5}

// 用三个线程分别打印出 a b c，然后让其循环打印 abcabcabc。
func GoroutinePritABC() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	current := 0 // 0=a, 1=b, 2=c

	for id := range 3 {
		go func(id int) {
			for {
				mu.Lock()
				for current != id { // 如果不是自己该打印的时候，就等待
					cond.Wait()
				}
				fmt.Print(string(rune('a' + id)))
				current = (current + 1) % 3 // 更新下一个该打印的字母
				// cond.Broadcast()            // 唤醒其他 goroutine
				mu.Unlock()
				time.Sleep(time.Millisecond * 100)
			}
		}(id)
	}
	select {} // 防止退出
}
