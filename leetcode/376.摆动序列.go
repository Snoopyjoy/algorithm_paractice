/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	up := 1
	down := 1
	for i:=1;i<len(nums);i++{
		if nums[i] > nums[i-1] {
			up = down+1
		}
		if nums[i-1] > nums[i] {
			down = up+1
		}
	}
	if up > down {
		return up
	}
	return down
}
// @lc code=end

