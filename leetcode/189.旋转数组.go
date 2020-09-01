/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	k = k % len(nums)
	if k == 0 {
		return
	}

	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for ; start < end; start, end = start+1, end-1 {
		nums[start], nums[end] = nums[end], nums[start]
	}
}

// @lc code=end

