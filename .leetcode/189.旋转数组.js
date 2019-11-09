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

    //最大公约数
    let mVal = measure( numLen, moveTimes );

    //外层循环次数为最大公约数
    const outerLoopTimes = mVal;
    //内层循环次数
    const innerLoopTimes = numLen / mVal;
    let temp;
    let temp1;
    let targetIndex;

    let oindex = 0;
    let index = 0;
    while( oindex < outerLoopTimes ){
        targetIndex = oindex;
        temp = nums[targetIndex];
        while( index < innerLoopTimes ){
            targetIndex = ( targetIndex + moveTimes ) % numLen;
            temp1 = nums[targetIndex];
            nums[targetIndex] = temp;
            temp = temp1;
            index++;
        }
        oindex++;
        index = 0;
    }
}

//求最大公约数
function measure( x,  y)  
{     
    let z = y;  
    while(x%y !== 0)  
    {  
        z = x%y;  
        x = y;  
        y = z;    
    }  
    return z;  
}

var rotate1 = function(nums, k) {
    const numLen = nums.length;
    const moveTimes = k % numLen;
    if( moveTimes === 0 || numLen <= 1) return;
    const orgNums = nums.concat([]); //暂存原数组
    nums.forEach( ( val, index )=>{
        nums[( index + moveTimes ) % numLen] = orgNums[index]; //替换移动后的新值
    } );
};

var rotate2 = function(nums, k) {
    const numLen = nums.length;
    const moveTimes = k % numLen;
    if( moveTimes === 0 || numLen <= 1) return;
    for (let index = 0; index < moveTimes; index++) {
        nums.unshift(nums.pop());
    }
}
// @lc code=end

