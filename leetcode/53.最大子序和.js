/*
 * @lc app=leetcode.cn id=53 lang=javascript
 *
 * [53] 最大子序和
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
    let sum = 0;
    let ans = nums[0];
    nums.forEach( num=>{
        if( sum > 0 ){
            sum += num;
        }else{
            sum = num;
        }
        ans = Math.max(sum,ans);
    } );
    return ans;
}

 //暴力法
var maxSubArray1 = function(nums) {
    let maxVal = Number.MIN_SAFE_INTEGER;
    for (let i = 0; i < nums.length; i++) {
        const val = nums[i];
        let sumVal = val;
        if( sumVal > maxVal ){
            maxVal = sumVal;
        }
        for (let j = i+1; j < nums.length; j++) {
            const bVal = nums[j];
            sumVal += bVal;
            if( sumVal > maxVal ){
                maxVal = sumVal;
            }
        }
    }
    return maxVal;
};

// @lc code=end


