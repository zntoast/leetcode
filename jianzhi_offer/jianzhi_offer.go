package jianzhioffer

import (
	"math"
)

//整数除法
func Divide(a int, b int) int {
	// 判断结果是否为正数 true正
	sign := 1
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}
	if (a > 0 && b > 0) || (a < 0 && b < 0) {
		sign = 1
	} else {
		sign = -1
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	ans := 0
	for i := 31; i >= 0; i-- {
		if (a>>i)-b >= 0 {
			a -= (b << i)
			ans += (1 << i)
		}
	}

	return ans * sign
}
