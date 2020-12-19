/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := -1
		for _, c := range coins {
			if i >= c && dp[i-c] > -1 {
				if min == -1 || dp[i-c]+1 < min {
					min = dp[i-c] + 1
				}
			}
		}
		dp[i] = min
	}
	return dp[amount]
}

// @lc code=end

