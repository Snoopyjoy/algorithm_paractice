/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	meetNode := findMeetNode(head)
	if meetNode == nil {
		return nil
	}
	for head != nil && meetNode != nil {
		if head == meetNode {
			return head
		}
		head = head.Next
		meetNode = meetNode.Next
	}
	return nil
}

func findMeetNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}

// @lc code=end

