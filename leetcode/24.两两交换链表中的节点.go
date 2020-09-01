/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := head.Next
	pre := &ListNode{
		Val:  0,
		Next: head,
	}
	for head != nil && head.Next != nil {
		head, pre.Next, pre, head.Next.Next, head.Next = head.Next.Next, head.Next, head, head, head.Next.Next
	}

	return ret
}

// @lc code=end

