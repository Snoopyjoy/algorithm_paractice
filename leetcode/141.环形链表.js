/*
 * @lc app=leetcode.cn id=141 lang=javascript
 *
 * [141] 环形链表
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
 * @return {boolean}
 */
var hasCycle = function(head) {
    if( head === null || head.next === null ) return false;
    let pointA = head;
    let pointB = head.next;
    let result = true;
    while( pointA !== pointB ){
        if( pointB === null || pointB.next === null ){
            result = false;
            break;
        }
        pointA = pointA.next;
        pointB = pointB.next.next;
    }
    return result;
};
// @lc code=end

