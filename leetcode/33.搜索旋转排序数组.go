package leetcode

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if goLeft(nums, target, left, right, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func goLeft(nums []int, target, left, right, mid int) bool {
	// 左边单调递增 且 target 小于 mid 大于 left
	con1 := nums[left] <= nums[mid] && target < nums[mid] && target >= nums[left]
	// 右边单调递增 且 target 小于 mid
	con2 := nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right])
	return con1 || con2
}

// @lc code=end
