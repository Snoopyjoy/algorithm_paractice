/*
 * @lc app=leetcode.cn id=206 lang=javascript
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function(head) {
    let temp = head;
    let tempNext = null;
    while( temp !== null ){
        let nextNode = temp.next;
        temp.next = tempNext;
        tempNext = temp;
        temp = nextNode;
    }

    return tempNext;
};
// @lc code=end

