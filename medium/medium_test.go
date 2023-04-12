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
		{[]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
			[]string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}},
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
