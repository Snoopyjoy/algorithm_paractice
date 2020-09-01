import "fmt"

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	if len(heights) < 1 {
		return 0
	}
	stack := make([]int, 0, len(heights))
	stack = append(stack, -1)
	area := 0
	for i := 0; i < len(heights); i++ {
		v := heights[i]
		for len(stack) > 1 && v < heights[top(stack)] {
			topI := top(stack)
			stack = stack[:len(stack)-1]
			preI := top(stack)
			tempA := (i - preI - 1) * heights[topI]
			if tempA > area {
				area = tempA
			}
		}
		stack = append(stack, i)
	}
	fmt.Println(stack)
	for len(stack) > 1 {
		topI := top(stack)
		stack = stack[:len(stack)-1]
		preI := top(stack)
		tempA := (len(heights) - preI - 1) * heights[topI]
		if tempA > area {
			area = tempA
		}
	}
	return area
}

func top(arr []int) int {
	return arr[len(arr)-1]
}

// @lc code=end

// 将问题转化为求单调递增柱状图中的最大矩形
// 采用单调递辅助增栈的方法
// 辅助栈中存入的是数组下标
// 顺序遍历数组
// 如果遍历到的元素比栈顶元素大直接入栈
// 如果遍历到的元素比栈顶元素小，则栈顶元素出栈并计算此时目前位置到栈第二位可形成区域的最大值，直到栈顶比自己小
// 求单调递增柱状图的最大矩形面积


