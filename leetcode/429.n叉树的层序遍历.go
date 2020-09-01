/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*Node{root}
	ret := [][]int{{root.Val}}
	for len(queue) > 0 {
		tempQ := []*Node{}
		items := []int{}
		for _, n := range queue {
			for _, v := range n.Children {
				tempQ = append(tempQ, v)
				items = append(items, v.Val)
			}
		}
		if len(items) > 0 {
			ret = append(ret, items)
		}
		queue = tempQ
	}
	return ret

}

// @lc code=end

