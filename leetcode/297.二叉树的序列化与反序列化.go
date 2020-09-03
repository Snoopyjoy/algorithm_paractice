import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
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

type Codec struct {
	sb strings.Builder
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	defer this.sb.Reset()
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n != nil {
			this.sb.WriteString(strconv.Itoa(n.Val))
			queue = append(queue, n.Left)
			queue = append(queue, n.Right)
		} else {
			this.sb.WriteString("null")
		}
		this.sb.WriteString(",")
	}

	return this.sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	if len(nodes) == 0 {
		return nil
	}
	if nodes[0] == "" || nodes[0] == "null" {
		return nil
	}
	z, _ := strconv.Atoi(nodes[0])
	ret := &TreeNode{
		Val: z,
	}
	layer := []*TreeNode{ret}
	i := 1
	for i < len(nodes) && len(layer) > 0 {
		node := layer[0]
		layer = layer[1:]
		left := nodes[i]
		right := nodes[i+1]
		if left != "null" && left != "" {
			v, _ := strconv.Atoi(left)
			node.Left = &TreeNode{
				Val: v,
			}
			layer = append(layer, node.Left)
		}
		if right != "null" && right != "" {
			v, _ := strconv.Atoi(right)
			node.Right = &TreeNode{
				Val: v,
			}
			layer = append(layer, node.Right)
		}
		i += 2
	}
	return ret
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

