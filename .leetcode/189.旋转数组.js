/*
 * @lc app=leetcode.cn id=189 lang=javascript
 *
 * [189] 旋转数组
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var rotate = function(nums, k) {
    const numLen = nums.length;
    const moveTimes = k % numLen;
    if( moveTimes === 0 || numLen <= 1) return;
    const orgNums = nums.concat([]);
    nums.forEach( ( val, index )=>{
        let targetIndex = ( index + moveTimes ) % numLen;
        nums[targetIndex] = orgNums[index];
    } );
};
// @lc code=end

