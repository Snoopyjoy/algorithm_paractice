/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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

var pre *TreeNode

func isValidBST(root *TreeNode) bool {
	pre = nil
	return isValidBSTHelper(root)
}

func isValidBSTHelper(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBSTHelper(root.Left) {
		return false
	}
	//fmt.Println(pre)
	if pre != nil && pre.Val >= root.Val {
		return false
	}
	pre = root

	return isValidBSTHelper(root.Right)
}

// @lc code=end

