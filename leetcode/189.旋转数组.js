/*
 * @lc app=leetcode.cn id=189 lang=javascript
 *
 * [189] 旋转数组
 * 
 */

//【每日一题】解题思路 将数组中的元素依次放到变换后的位置 索引为index的元素变换后应该放到 ( index + k ) % nums.length 的位置。
// 空间复杂度为O(1)，因此遍历时内部不可以声明新的变量，但可以在循环外声明临时变量来存储交换时的中间值
// 遍历时的处理为 位置index的元素放到index2 =( index + k ) % nums.length 的位置 index2的元素放到 index3 =( index2 + k ) % nums.length的位置。以此类推
// 当数组中所有元素都遍历完一次时变换完成
// 需要解决的问题是如何做到把所有元素都遍历一次，如 nums为[1,2,3,4]，k为2的情况下会发生0->2->0->2这样的索引元素交换的情况
// 经过观察，可以先求出数组长度和k的最大公约数mVal，使用双层循环的方式即可遍历所有元素，其中外层循环的次数为最大公约数mVal，内存循环的次数为(数组长度/mVal)
// 时间复杂度为O(n) n为数组长度 空间复杂度为常量O(1)

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

