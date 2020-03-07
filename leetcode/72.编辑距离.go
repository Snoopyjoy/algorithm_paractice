package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	// ""|word1
	// ""|word2
	word1Arr := strings.Split(word1, "")
	word2Arr := strings.Split(word2, "")
	dp := make([][]int, len(word1Arr)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2Arr)+1)
		dp[i][0] = i
	}

	for i := range dp[0] {
		dp[0][i] = i
	}
	for i := 1; i <= len(word1Arr); i++ {
		for j := 1; j <= len(word2Arr); j++ {
			if word1Arr[i-1] == word2Arr[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[len(word1Arr)][len(word2Arr)]
}

func min(a, b, c int) int {
	return min2(min2(a, b), c)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
