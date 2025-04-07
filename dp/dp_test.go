package dp

import (
	"fmt"
	"testing"
)

// 最长公共子序列
func TestLongestCommonSubsequence(t *testing.T) {
	text1 := "accde"
	text2 := "ace"
	i := LongestCommonSubsequence(text1, text2)
	fmt.Printf("i: %v\n", i)
}

func TestUniquePaths(t *testing.T) {
	i := UniquePaths(3, 2)
	fmt.Printf("i: %v\n", i)
}

func TestUniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := make([][]int, 3)
	for k := range obstacleGrid {
		obstacleGrid[k] = make([]int, 3)
	}
	obstacleGrid[1][1] = 1
	i := UniquePathsWithObstacles(obstacleGrid)
	fmt.Printf("i: %v\n", i)
}

func TestMinPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {1, 2, 1}}
	i := MinPathSum(grid)
	fmt.Printf("i: %v\n", i)
}

func TestMinimumTotal(t *testing.T) {
	grid := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	i := MinimumTotal(grid)
	fmt.Printf("i: %v\n", i)
}

func TestMinfallingpathsumI(t *testing.T) {
	grid := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	i := MinFallingPathSumI(grid)
	fmt.Printf("i: %v\n", i)
}

func TestMinFallingPathSumII(t *testing.T) {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	i := MinFallingPathSumII(grid)
	fmt.Printf("i: %v\n", i)
}

func TestRob(t *testing.T) {
	date := []struct {
		nums []int
	}{
		{[]int{1, 2, 3, 1}},
		{[]int{2, 7, 9, 3, 1}},
		{[]int{2, 1, 1, 2}}, //
	}
	for k, data := range date {
		result := Rob(data.nums)
		fmt.Printf("第%d次测试的结果 : %v\n", k+1, result)
	}
}
