/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	ret, pre := 3, 2
	for i := 3; i < n; i++ {
		ret, pre = pre+ret, ret
	}
	return ret
}

// @lc code=end

