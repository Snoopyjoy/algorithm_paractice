/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	if len(coins) == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, c := range coins {
		for i := 1; i <= amount; i++ {
			if i >= c {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	return dp[amount]
}

// @lc code=end

