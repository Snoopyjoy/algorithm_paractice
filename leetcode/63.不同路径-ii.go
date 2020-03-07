package leetcode

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	colBocked := false
	for i := range dp {
		dp[i] = make([]int, n)
		if !colBocked {
			if obstacleGrid[i][0] == 1 {
				colBocked = true
			} else {
				dp[i][0] = 1
			}
		}
		if i == 0 {
			for j := range dp[0] {
				if obstacleGrid[0][j] != 1 {
					dp[0][j] = 1
				} else {
					break
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

// @lc code=end
