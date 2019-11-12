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
    const stack = [ [-1, -1] ];      //高度,索引
    let curBottom = 0;
    let result = 0;
    height.forEach( (val, index)=>{
        const lastEle = stack[stack.length-1];
        const lastVal = lastEle[0];

        console.log( "---------------------------" );
        console.log( "lastVal", lastVal);
        console.log( "curVal", val);

        stack.push(  [val, index] );
        if( val >= lastVal ){   //检查是否有闭合区间并收集
            result += collectStack( stack, curBottom );
            curBottom = lastVal;
        }
    } );
    return result;
};

function collectStack( stack, bottom ){
    let result = 0;
    const wall = stack[stack.length-1];
    const wallVal = wall[0];
    const wallIndex = wall[1];

    const startWall = stack[1];
    if(!startWall) return 0;

    const startWallVal = startWall[0];
    const startWallIndex = startWall[1];

    let stopIndex = -1;
    
    let startIndex = 1;
    //查找起始蓄水起始索引
    for (let i = stack.length - 2; i >=0; i--) {

    }

    for (let i = stack.length - 2; i >=0; i--) {
        const stackEle = stack[i];
        const val = stackEle[0];
        const index = stackEle[1];

        if( i === startWallIndex ){

            let height = wallVal - val - bottom;
            height = height < 0 ? 0 : height;
            result += height * (wallIndex-index);
            stopIndex = i;
            break;
        }

        if( val < wallVal ){
            let height = wallVal - val - bottom;
            height = height < 0 ? 0 : height;
            result += height * (wallIndex-index);
        }else{
            stopIndex = i;
            break;
        }
    }
    console.log("stopIndex",stopIndex)
    if( stopIndex === -1 ) return 0;
    stack.length = stopIndex+1;
    console.log(stack)
    return result;
}
// @lc code=end

//trap([0,1,0,2,1,0,1,3,2,1,2,1]);

const result = trap([0,1,0,2]);
console.log(result, result);