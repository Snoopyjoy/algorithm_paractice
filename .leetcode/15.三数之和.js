/*
 * @lc app=leetcode.cn id=15 lang=javascript
 *
 * [15] 三数之和
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
    if( nums.length < 3 ) return [];
    let result = [];
    nums.sort((a,b)=>a-b);
    //console.log(nums);
    if( nums[0] > 0 ) return [];
    let i = 0;
    let left = 1;
    let right = nums.length - 1;
    while( i < nums.length - 2 && nums[i] <= 0){
        while(left < right){
            let sum = nums[i] + nums[left] + nums[right];
            //console.log(i,left ,right, sum);
            if( sum === 0 ){
                result.push(
                    [ nums[i], nums[left], nums[right] ]
                );
                while( left<right && nums[left+1] === nums[left] ) left++;
                while( left<right && nums[right-1] === nums[right] ) right--;
                left++;
                right--;
            }else if( sum < 0 ){
                left++;
            }else if( sum > 0 ){
                right--;
            }
        }
        while( i < nums.length - 2 && nums[i] === nums[i+1] ) i++;
        i++;
        left = i + 1;
        right = nums.length - 1;
    }
    return result;
};
// @lc code=end

//threeSum([-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]);

