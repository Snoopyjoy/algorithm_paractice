/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tempQueue := []*TreeNode{}
		if queue[0].Left != nil {
			tempQueue = append(tempQueue, queue[0].Left)
		}
		if queue[0].Right != nil {
			tempQueue = append(tempQueue, queue[0].Right)
		}
		max := queue[0].Val
		for i := 1; i < len(queue); i++ {
			itm := queue[i]
			if itm.Val > max {
				max = itm.Val
			}
			if itm.Left != nil {
				tempQueue = append(tempQueue, itm.Left)
			}
			if itm.Right != nil {
				tempQueue = append(tempQueue, itm.Right)
			}
		}
		ret = append(ret, max)
		queue = tempQueue
	}

	return ret
}

// @lc code=end

