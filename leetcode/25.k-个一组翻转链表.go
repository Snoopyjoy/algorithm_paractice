/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	// 上一段的尾部
	pre := dummy
	cur := dummy
	for cur != nil && pre != nil {
		for i := 0; i < k && cur != nil; i++ {
			cur = cur.Next
		}
		if cur == nil {
			break
		}
		temp := cur.Next
		start := pre.Next
		// 截断
		cur.Next = nil
		l1 := reverseList(start)
		start.Next = temp
		pre.Next = l1

		pre = start
		cur = start
	}

	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

// @lc code=end

