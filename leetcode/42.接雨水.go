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
	stack := []int{-1}
	ret := 0
	for i, h := range height {
		for len(stack) > 1 && h >= height[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) < 2 {
				break
			}
			preIndex := stack[len(stack)-1]
			preH := height[preIndex]
			if h > preH {
				h = preH
			}
			ret += (i - preIndex - 1) * (h - height[topIndex])
		}
		stack = append(stack, i)
	}

	return ret
}

// @lc code=end

