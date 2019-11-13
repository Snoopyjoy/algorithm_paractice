/*
 * @lc app=leetcode.cn id=42 lang=javascript
 *
 * [42] 接雨水
 */

// @lc code=start
/**
 * @param {number[]} height
 * @return {number}
 */
var trap = function(height) {
    const stack = [];      //索引
    let result = 0;
    height.forEach( (val, index)=>{
        while( stack.length > 0 && val > height[ stack[stack.length-1] ] ){
            const tail = stack[stack.length-1];
            stack.pop();
            if( stack.length === 0 ){
                break;
            }
            const prev = stack[stack.length-1];
            const steps = index - prev - 1;
            const prevH = height[prev];
            const h = Math.min( prevH, val );
            const bottom = height[tail];

            result += (h - bottom) * steps;
        }
        stack.push(index);
    } );
    return result;
};
// @lc code=end

//const result = trap([0,7,1,4,6]);
//const result = trap([5,2,1,2,1,5]);
//const result = trap([0,1,0,2,1,0,1,3,2,1,2,1]);

//const result = trap([0,1,0,2]);
//console.log("result", result);