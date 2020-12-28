import "fmt"

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stk := make([]int, 0, len(heights)+1)
	stk = append(stk, -1)

	maxArea := 0
	for i, h := range heights {
		for len(stk) > 1 && h <= heights[stk[len(stk)-1]] {
			hi := heights[stk[len(stk)-1]]
			stk = stk[:len(stk)-1]
			preI := stk[len(stk)-1]

			area := (i - preI - 1) * hi
			if area > maxArea {
				maxArea = area
			}
		}
		stk = append(stk, i)
	}

	for i := 1; i < len(stk); i++ {
		preI := stk[i-1]
		hi := heights[stk[i]]
		area := (len(heights) - preI - 1) * hi
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// @lc code=end

// 将问题转化为求单调递增柱状图中的最大矩形
// 采用单调递辅助增栈的方法
// 辅助栈中存入的是数组下标
// 顺序遍历数组
// 如果遍历到的元素比栈顶元素大直接入栈
// 如果遍历到的元素比栈顶元素小，则栈顶元素出栈并计算此时目前位置到栈第二位可形成区域的最大值，直到栈顶比自己小
// 求单调递增柱状图的最大矩形面积


