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
    const numDic = new Map();
    let tempIndexs = [];
    let length = 0;
    for (let index = 0; index < nums.length; index++) {
        const numVal = nums[index];
        if( numDic.has( numVal ) ) {
            tempIndexs.push( index );
        }else{
            numDic.set( numVal, true );
            length++;
            let opIndex = tempIndexs.shift();
            if( opIndex > 0 ){
                nums[opIndex] = numVal;
                tempIndexs.push( index );
            }
        }
    }
    return length;
};
// @lc code=end

