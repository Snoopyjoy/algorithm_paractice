/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	stack := make([]int, 0, len(height)+1)
	stack = append(stack, -1)
	ret := 0
	for i, v := range height {
		for len(stack) > 1 && v > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) < 2 {
				break
			}
			sec := stack[len(stack)-1]
			h := v
			if v > height[sec] {
				h = height[sec]
			}
			ret += (h - height[top]) * (i - sec - 1)
		}

		stack = append(stack, i)
	}
	return ret
}

// @lc code=end

