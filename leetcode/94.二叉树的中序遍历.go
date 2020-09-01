package leetcode

import "container/list"

// TreeNode Val Left Right
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	// ret := []int{}
	// dfs(root, &ret)
	// return ret
	return stackWay(root)
}

// 栈的方式遍历
func stackWay(root *TreeNode) []int {
	stack := list.New()
	result := []int{}
	curNode := root
	for stack.Len() > 0 || curNode != nil {
		// 左节点入栈
		for curNode != nil {
			stack.PushBack(curNode)
			curNode = curNode.Left
		}

		// 左节点出栈并放入结果 下一轮遍历该节点的右子树（叶子节右子树为空）
		if stack.Len() != 0 {
			curNode = (stack.Back().Value).(*TreeNode)
			result = append(result, curNode.Val)
			stack.Remove(stack.Back())
			curNode = curNode.Right
		}
	}
	return result
}

// 递归遍历
func dfs(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, result)
	*result = append(*result, node.Val)
	dfs(node.Right, result)
}

// @lc code=end
