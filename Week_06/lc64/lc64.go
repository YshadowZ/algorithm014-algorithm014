package lc64

import "fmt"

// 左上到右下，第一行每一个位置只能是去左边过来的，第一列每一个位置只能是上面下来的
func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, columns)
	}
	dp[0][0] = grid[0][0]
	// 第一行
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// 第一列
	for j := 1; j < columns; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	fmt.Println(dp)
	return dp[rows-1][columns-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
