package leetcode

import "container/list"

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	ret := []int{}
	preDfs(root, &ret)
	return ret
	//return preStackWay(root)
}

func preStackWay(root *TreeNode) []int {
	stack := list.New()
	result := []int{}
	curNode := root
	for {
		if curNode == nil && stack.Len() == 0 {
			break
		}
		for {
			if curNode == nil {
				break
			}
			result = append(result, curNode.Val)
			stack.PushBack(curNode)
			curNode = curNode.Left
		}

		if stack.Len() > 0 {
			curNode = (stack.Back().Value).(*TreeNode)
			stack.Remove(stack.Back())
			curNode = curNode.Right
		}
	}

	return result
}

// 递归
func preDfs(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	preDfs(node.Left, result)
	preDfs(node.Right, result)
	return
}

// @lc code=end
