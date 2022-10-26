package dp

import (
	"math"
	"sort"
)

/*
	动态规划算法题
	题解
	from: https://github.com/zntoast
*/

/*
剑指 Offer II 095. 最长公共子序列
*/
func LongestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, len1+1)
	for k := range dp {
		dp[k] = make([]int, len2+1)
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, t1 := range text1 {
		for j, t2 := range text2 {
			if t1 == t2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[len1][len2]
}

/*
LC 62.不同路径I
*/
func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else if i > 0 {
				dp[i][j] = dp[i-1][j]
			} else if j > 0 {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

/*
LC 63.不同路径II
*/
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i > 0 && j > 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else if i > 0 {
				dp[i][j] = dp[i-1][j]
			} else if j > 0 {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

/*
LC 64.最小路径和
*/
func MinPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			} else if i > 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else if j > 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

/*
LC 120.三角形最小路径和
*/
func MinimumTotal(triangle [][]int) int {
	m := len(triangle)
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, m)
	}
	dp[0][0] = triangle[0][0]
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	//最后一排有多少个数就循环多少次
	for i := 1; i < m; i++ {
		for j := 0; j < i+1; j++ {
			dp[i][j] = math.MaxInt32
			val := triangle[i][j]
			if j != 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+val)
			}
			if i != j {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+val)
			}
		}
	}
	sort.Ints(dp[m-1])
	return dp[m-1][0]
}

/*
LC 931.下降路径最小和I
*/
func MinFallingPathSumI(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	copy(dp[0], matrix[0])
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	//最后一排有多少个数就循环多少次
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
			val := matrix[i][j]
			if j != 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+val)
			}
			if j != m-1 {
				dp[i][j] = min(dp[i][j], dp[i-1][j+1]+val)
			}
			dp[i][j] = min(dp[i][j], dp[i-1][j]+val)
		}
	}
	sort.Ints(dp[m-1])
	return dp[m-1][0]
}

/*
LC 1289.下降路径最小和 II
*/
func MinFallingPathSumII(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	copy(dp[0], grid[0])
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	//最后一排有多少个数就循环多少次
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = math.MaxInt32
			val := grid[i][j]
			for k := 0; k < n; k++ {
				if j == k {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k]+val)
			}
		}
	}
	sort.Ints(dp[m-1])
	return dp[m-1][0]
}

/*
LC 1575. 统计所有可行路径
*/
func CountRoutes(locations []int, start int, finish int, fuel int) int {
	return 0
}
