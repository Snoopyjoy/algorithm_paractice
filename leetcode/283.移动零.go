/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	slow := 0
	for i, v := range nums {
		nums[i], nums[slow] = nums[slow], nums[i]
		if v != 0 {
			slow++
		}
	}
}

// @lc code=end
