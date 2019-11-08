/*
 * @lc app=leetcode.cn id=26 lang=javascript
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    if( nums.length === 0 ) return 0;

    let cursor = 0;
    nums.forEach( ( num )=>{
        if( nums[cursor] != num){
            nums[++cursor] = num;
        }
    } )

    return cursor+1;
};
// @lc code=end

