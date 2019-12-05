/*
 * @lc app=leetcode.cn id=33 lang=javascript
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
    // 先转为有序数组
    const revPos = findRevPos(nums, 0, nums.length);
    const right = nums.splice(0,revPos+1);
    nums = nums.concat(right);
    // 二分查找
    const newPos = findPos( nums, target, 0, nums.length );
    const orgRightLength = nums.length - revPos - 1;
    if( revPos > -1 && newPos > -1){
        if ( newPos > orgRightLength - 1 ) {
            return newPos - orgRightLength;
        }else{
            return newPos + revPos + 1;
        }
    }
    return newPos;
    
};

function findPos(nums, target, left, right) {
    if( left >= right ) return -1;
    let middlePos = left + (((right - left)/2)>>0);
    let middleVal = nums[middlePos];
    if( middleVal === target ) return middlePos;
    if( left === middlePos ) return -1;
    if (middleVal > target) {
        return findPos(nums, target, left, middlePos);
    }
    return findPos(nums, target, middlePos, right);
}

function findRevPos(nums, left, right){
    if( left >= right ) return -1;
    let middlePos = left + (((right - left)/2)>>0);
    let middleVal = nums[middlePos];
    if( middlePos < nums.length - 1 && middleVal > nums[middlePos+1] ) return middlePos;
    if( middlePos > 0 && middleVal < nums[middlePos-1] ) return middlePos-1;
    if( middleVal > nums[right - 1] ) return findRevPos(nums, middlePos, right);
    if( middleVal < nums[left] ) return findRevPos(nums, left, middlePos);
    return -1;
}

// @lc code=end

