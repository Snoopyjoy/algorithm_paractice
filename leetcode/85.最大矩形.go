/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	lefts := make([][]int, len(matrix))

	for i, v := range matrix {
		line := make([]int, len(v))
		n := 0
		for j, jv := range v {
			if jv == '0' {
				n = 0
			} else {
				n++
			}
			line[j] = n
		}
		lefts[i] = line
	}

	maxArea := 0
	for j := 0; j < len(lefts[0]); j++ {
		stack := make([]int, len(lefts)+1)
		stack[0] = -1
		for i, line := range lefts {
			for len(stack) > 1 && line[j] <= lefts[stack[len(stack)-1]][j] {
				hi := lefts[stack[len(stack)-1]][j]
				stack = stack[:len(stack)-1]
				preI := stack[len(stack)-1]
				area := hi * (i - preI - 1)
				if area > maxArea {
					maxArea = area
				}
			}
			stack = append(stack, i)
		}
		for i := 1; i < len(stack); i++ {
			preI := stack[i-1]
			hi := lefts[stack[i]][j]
			area := hi * (len(lefts) - preI - 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// @lc code=end

// 统计i左边连续1的个数
// 转化为柱状图
// 求柱状图的最大矩形
// 柱状体简化为单调递增柱状图
// 求单调递增柱状图的最大矩形
