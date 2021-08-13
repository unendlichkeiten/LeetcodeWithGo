package _064_minimum_path_sum

func minPathSum(grid [][]int) int {
	// 状态转移方程
	// dp[i][j] = min(dp[i][j+1], dp[i+1][j]) + grid[i][j]

	// 初始化 dp
	dp := [][]int{}
	for i := 0; i < len(grid); i++ {
		tmp := make([]int, len(grid[0]))
		dp = append(dp, tmp)
	}

	i, j := 0, 0
	for i = 0; i < len(grid); i++ {
		for j = 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}

	return dp[i-1][j-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
