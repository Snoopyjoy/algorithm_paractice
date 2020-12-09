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
	d := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := d
	pre := d

	for cur != nil && pre != nil {
		for i := 0; i < k && cur != nil; i++ {
			cur = cur.Next
		}
		if cur == nil {
			break
		}
		tmp := cur.Next
		cur.Next = nil
		start := pre.Next
		n := reverseList(start)
		pre.Next = n
		start.Next = tmp
		pre = start
		cur = start
	}
	return d.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		cur, pre, cur.Next = cur.Next, cur, pre
	}
	return pre
}

// @lc code=end

