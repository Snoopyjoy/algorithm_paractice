/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{}
	dfsHelper(root, &ret, 0)
	return ret
}

func dfsHelper(node *TreeNode, ret *[]int, depth int) {
	if node == nil {
		return
	}
	if len(*ret) == depth {
		*ret = append(*ret, node.Val)
	}
	depth++
	dfsHelper(node.Right, ret, depth)
	dfsHelper(node.Left, ret, depth)
}

// @lc code=end

// 广度优先遍历
func rightSideViewBFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tempQueue := []*TreeNode{}
		for i, itm := range queue {
			if i == 0 {
				ret = append(ret, itm.Val)
			}
			if itm.Right != nil {
				tempQueue = append(tempQueue, itm.Right)
			}
			if itm.Left != nil {
				tempQueue = append(tempQueue, itm.Left)
			}
		}
		queue = tempQueue
	}
	return ret
}
