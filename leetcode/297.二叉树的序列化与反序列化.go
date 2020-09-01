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
	//this.sb.WriteString("[")
	this.serializeHelper(root)
	//this.sb.WriteString("]")
	return this.sb.String()
}
func (this *Codec) serializeHelper(root *TreeNode) {
	if root == nil {
		this.sb.WriteString("null,")
		return
	}
	this.sb.WriteString(strconv.FormatInt(int64(root.Val), 10))
	this.sb.WriteString(",")
	this.serializeHelper(root.Left)
	this.serializeHelper(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	allNodes := strings.Split(data, ",")
	if len(allNodes) == 0 {
		return nil
	}
	// // fmt.Println("allNodes", allNodes)
	counter := 0
	return this.deserializeHelper(allNodes, &counter)
}
func (this *Codec) deserializeHelper(allNodes []string, counter *int) *TreeNode {
	if *counter >= len(allNodes) {
		return nil
	}
	// fmt.Println(*counter)
	if allNodes[*counter] == "null" || allNodes[*counter] == "" {
		*counter = *counter + 1
		return nil
	}
	v, _ := strconv.ParseInt(allNodes[*counter], 10, 64)
	ret := &TreeNode{
		Val: int(v),
	}
	*counter = *counter + 1
	ret.Left = this.deserializeHelper(allNodes, counter)
	ret.Right = this.deserializeHelper(allNodes, counter)
	return ret
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
// @lc code=end

