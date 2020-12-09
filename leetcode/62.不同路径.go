package leetcode_62

import "fmt"

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

// 一维数组
// 数组对应位置的原始值就是格子正上方格子的值
// 左边格子的值加上正上方格子的值就是当前格子的值

// @lc code=end
