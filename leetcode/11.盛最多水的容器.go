/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	ret, left, right := 0, 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			ret = max(ret, (right-left)*min(height[left], height[right]))
			left++
		} else {
			ret = max(ret, (right-left)*min(height[left], height[right]))
			right--
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

