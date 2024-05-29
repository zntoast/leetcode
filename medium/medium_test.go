package medium

import (
	"fmt"
	"testing"
)

func TestAlertNames(t *testing.T) {
	date := []struct {
		keyName []string
		keyTime []string
	}{
		{
			[]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
			[]string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"},
		},
	}
	for k, d := range date {
		count := AlertNames(d.keyName, d.keyTime)
		fmt.Printf("第%d次测试的结果 : %v\n", k+1, count)
	}
}

func TestPartitionLabels(t *testing.T) {
	date := []struct {
		s string
	}{
		{"ababcbacadefegdehijhklij"},
		{"eccbbbbdec"},
	}
	for k, d := range date {
		result := PartitionLabels(d.s)
		fmt.Printf("第%d次测试的结果 : 分割的数组%v\n", k+1, result)
	}
}

func TestMonotoneIncreasingDigits(t *testing.T) {
	date := []struct {
		s int
	}{
		{235491},
		{192},
		{1234},
	}
	for k, d := range date {
		result := MonotoneIncreasingDigits(d.s)
		fmt.Printf("第%d次测试的结果 : 比 %d小的最大单调整数%v\n", k+1, d.s, result)
	}
}

func TestConvert(t *testing.T) {
	date := []struct {
		s       string
		numRows int
	}{
		{"PAYPALISHIRING", 3},
		{"A", 1},
	}
	for k, data := range date {
		result := Convert(data.s, data.numRows)
		fmt.Printf("第%d次测试的结果 : 输出字符：%v\n", k+1, result)
	}
}

func TestDailyTemperatures(t *testing.T) {
	date := []struct {
		temperatures []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}},
		{[]int{30, 40, 50, 60}},
		{[]int{30, 60, 90}},
	}
	for k, data := range date {
		result := DailyTemperatures(data.temperatures)
		fmt.Printf("第%d次测试的结果 : 结果数组：%v\n", k+1, result)
	}
}

func TestIntToRoman(t *testing.T) {
	date := []struct {
		num int
	}{
		{500},
		{300},
		{58},
		{1994},
	}
	for k, data := range date {
		result := IntToRoman(data.num)
		fmt.Printf("第%d次测试的结果 : 结果数组：%v\n", k+1, result)
	}
}

func TestGeneratetarenthesis(t *testing.T) {
	date := []struct {
		num int
	}{
		// {4},
		// {1},
		{3},
	}
	for k, data := range date {
		result := GenerateParenthesis(data.num)
		fmt.Printf("第%d次测试的结果 : 结果数组：%v\n", k+1, result)
	}
}

func TestSwapPairs(t *testing.T) {
	n3 := &ListNode{4, nil}
	n2 := &ListNode{3, n3}
	n1 := &ListNode{2, n2}
	head := &ListNode{1, n1}
	SwapPairs(head)
	fmt.Printf("head: %v\n", head)
}

func TestDivide(t *testing.T) {
	date := []struct {
		dividend int
		divisor  int
	}{
		{3, -1},
		{3, 1},
		{-3, 1},
		{10, 3},
	}
	for k, data := range date {
		result := Divide(data.dividend, data.divisor)
		fmt.Printf("第%d次测试的结果 : 结果数组：%v\n", k+1, result)
	}
}

func TestNextPermutation(t *testing.T) {
	date := []struct {
		nums []int
	}{
		{[]int{1, 2, 3}},
	}
	for k, data := range date {
		NextPermutation(data.nums)
		fmt.Printf("第%d次测试的结果 : 结果数组：%v\n", k+1, data.nums)
	}
}

func TestPermutation(t *testing.T) {
	date := []struct {
		S string
	}{
		{"abd"},
	}
	for k, data := range date {
		result := Permutation(data.S)
		fmt.Printf("第%d次测试的结果 : 结果数组：%v\n", k+1, result)
	}
}

func TestCountAndSay(t *testing.T) {
	date := []struct {
		n int
	}{
		{3},
		{1},
		{4},
		{5},
		{6},
	}
	for k, data := range date {
		result := CountAndSay(data.n)
		fmt.Printf("第%d次测试的结果 : %d 对应的值为 %s\n", k+1, data.n, result)
	}
}

func TestCanJump(t *testing.T) {
	date := []struct {
		nums []int
	}{
		// {[]int{2, 3, 1, 1, 4}},
		// {[]int{3, 2, 1, 0, 4}},
		{[]int{1, 2}},
	}
	for k, data := range date {
		result := CanJump(data.nums)
		fmt.Printf("第%d次测试的结果  是否能到达最后一个下标为%v\n", k+1, result)
	}
}

func TestCombine(t *testing.T) {
	date := []struct {
		n int
		k int
	}{
		{4, 2},
		{1, 1},
	}
	for k, data := range date {
		result := Combine(data.n, data.k)
		fmt.Printf("第%d次测试的结果  是否能到达最后一个下标为%v\n", k+1, result)
	}
}

func TestJump(t *testing.T) {
	date := []struct {
		nums []int
	}{
		{[]int{2, 3, 1, 1, 4}},
		{[]int{2, 3, 0, 1, 4}},
	}
	for k, data := range date {
		result := Jump(data.nums)
		fmt.Printf("第%d次测试的结果  是否能到达最后一个下标为%v\n", k+1, result)
	}
}

func TestMultiply(t *testing.T) {
	date := []struct {
		num1 string
		num2 string
	}{
		{"2", "3"},
		{"123", "456"},
	}
	for k, data := range date {
		result := Multiply(data.num1, data.num2)
		fmt.Printf("第%d次测试的结果  %v * %v乘积%v\n", k+1, data.num1, data.num2, result)
	}
}

func TestSingleNumber(t *testing.T) {
	date := []struct {
		num []int
	}{
		{[]int{1, 2, 1, 3, 2, 5}},
		{[]int{-1, 0}},
		{[]int{0, 1}},
	}
	for k, data := range date {
		result := SingleNumber(data.num)
		fmt.Printf("第%d次测试的结果  %v \n", k+1, result)
	}
}

func TestCombinationSum(t *testing.T) {
	datas := []struct {
		candidates []int
		target     int
	}{
		{[]int{2, 3, 6, 7}, 7},
		{[]int{2, 3, 5}, 8},
	}
	for k, data := range datas {
		result := CombinationSum(data.candidates, data.target)
		fmt.Printf("第%d次测试的结果  %v \n", k+1, result)
	}
}

func TestNumberOfWays(t *testing.T) {
	datas := []struct {
		StartPos int
		EndPos   int
		K        int
	}{
		{1, 2, 3},
		{2, 5, 10},
	}
	for k, data := range datas {
		result := NumberOfWays(data.StartPos, data.EndPos, data.K)
		fmt.Printf("第%d次测试的结果  %v \n", k+1, result)
	}
}

func TestMaxArrayValue(t *testing.T) {
	datas := []struct {
		nums []int
	}{
		{[]int{5, 7, 9, 3}},
	}
	for k, data := range datas {
		result := MaxArrayValue(data.nums)
		fmt.Printf("第%d次测试的结果  %v \n", k+1, result)
	}
}

func TestLongestEqualSubarray(t *testing.T) {
	datas := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 3, 2, 3, 1, 3}, 3},
		// {[]int{}, 12447},
	}
	for k, data := range datas {
		result := LongestEqualSubarray(data.nums, data.k)
		t.Logf("第%d次测试的结果  %v \n", k+1, result)
	}
}

func TestMaximumLength(t *testing.T) {
	datas := []struct {
		s string
	}{
		{"aaaa"},
		{"abcdef"},
		{"abcaba"},
		{"azzyyyyyyy"},
	}
	for k, data := range datas {
		result := MaximumLength(data.s)
		t.Logf("第%d次测试的结果  %v \n", k+1, result)
	}
}
