package leetcode

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	preRob := nums[0]
	preNoRob := 0
	for i := 1; i < len(nums); i++ {
		curRob := preNoRob + nums[i]
		preNoRob = max(preNoRob, preRob)
		preRob = curRob
	}
	return max(preRob, preNoRob)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
