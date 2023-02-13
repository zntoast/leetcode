package medium

import (
	"sort"
)

/*
LC 763. 划分字母区间
*/
func PartitionLabels(s string) []int {
	temp := map[byte]int{}
	for i := 0; i < len(s); i++ {
		temp[s[i]] = i
	}
	result := make([]int, 0)
	left := 0
	right := 0
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for j := 0; j < len(s); j++ {
		right = max(right, temp[s[j]])
		if right == j {
			result = append(result, right-left+1)
			left = right + 1
		}
	}
	return result
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
