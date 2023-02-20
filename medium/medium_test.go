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
