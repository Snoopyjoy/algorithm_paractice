package leetcode.105

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	rootInorderIndex := -1
	for i, v := range inorder {
		if v == rootVal {
			rootInorderIndex = i
			break
		}
	}
	leftInorder := inorder[:rootInorderIndex]
	rightInorder := inorder[rootInorderIndex+1:]

	var rightPreorder []int
	var leftPreorder []int
	if len(leftInorder) > 0 {
		leftPreorder = preorder[1:rootInorderIndex+1]
	}
	rightPreorder = preorder[rootInorderIndex+1:]
	ret := &TreeNode{
		Val: rootVal,
		Left: buildTree(leftPreorder, leftInorder),
		Right: buildTree(rightPreorder, rightInorder),
	}

	return ret
}
// @lc code=end
