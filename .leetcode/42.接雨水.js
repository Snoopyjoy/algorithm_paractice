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
    let bottomInfo = {
        val: 0,
        index: -1
    };
    let result = 0;
    height.forEach( (val, index)=>{
        const lastEle = stack[stack.length-1];

        console.log( "---------------------------" );

        stack.push(  [val, index] );
        console.log("stackStart", stack);
        if( lastEle ){   //检查是否有闭合区间并收集雨水

            const lastVal = lastEle[0];
            const lastIndex = lastEle[1];
            const bottomVal = bottomInfo.val;

            if( val > lastVal || ( val > bottomVal && (index - lastIndex) > 1 ) ){

                const addWater = collectStack( stack, bottomInfo );
                console.log("addResult", addWater)
                result += addWater;
                console.log("result", result);
            
            }
        }

        console.log( stack );
    } );
    return result;
};

function collectStack( stack, bottomInfo ){
    let result = 0;
    const endStackIndex = stack.length-1;
    const endWall = stack[endStackIndex];
    const endWallVal = endWall[0];
    const endWallIndex = endWall[1];

    let startStackIndex = 0;
    let startWallVal = -1;

    let bottomVal = bottomInfo.val;
    let bottomIndex = bottomInfo.index;
    console.log("bottomVal",bottomVal);
    console.log("bottomIndex",bottomIndex);

    for (let i = stack.length - 2; i >=0; i--) {
        const eleData = stack[i];
        const val = eleData[0];
        if( val >= endWallVal ){
            startStackIndex = i;
            break;
        }
    }

    if( stack[startStackIndex] ){
        startWallVal = stack[startStackIndex][0];
        startWallIndex = stack[startStackIndex][1];
    }else{
        return result;
    }

    console.log("startWallVal",startWallVal  );
    console.log("startStackIndex",startStackIndex  );

    let wallHeight = 0;
    if( startWallVal > endWallVal ){
        wallHeight = endWallVal;
        bottomInfo.val = wallHeight;
        bottomInfo.index = endWallIndex;
    }else if(startWallVal < endWallVal){
        wallHeight = startWallVal;
        bottomInfo.val = 0;
        bottomInfo.index = -1;
    }else{
        wallHeight = startWallVal;
        if( startStackIndex === 0){
            bottomInfo.val = 0;
        }else{
            bottomInfo.val = wallHeight;
            bottomInfo.index = endWallIndex;
        }
    }

    console.log("wallHeight",wallHeight  );

    //计算区间蓄水
    for (let j = startStackIndex; j < stack.length-1; j++) {
        const eleInfo = stack[j];
        const eleVal = eleInfo[0];
        const eleIndex = eleInfo[1];

        const nextEleInfo = stack[j+1];
        const nextEleVal = nextEleInfo[0];
        const nextEleIndex = nextEleInfo[1];

        //距离
        const steps = nextEleIndex - eleIndex;
        
        const middleWater = (steps - 1) * (wallHeight - bottomVal);
        
        if( nextEleIndex > bottomIndex ){
            bottomVal = 0;
            bottomIndex = -1;
        }
        let height = Math.min( nextEleVal, wallHeight );
        height = Math.max( height, bottomVal );
        const topWater = wallHeight - height;
        let water = middleWater + topWater;
        water = water < 0 ? 0:water;
        result += water;
    }
    if(startWallVal > endWallVal){
        //除了开始位置留下
        stack.splice(startStackIndex+1);
    }else if( startWallVal === endWallVal && startStackIndex > 0){
        stack.splice(startStackIndex);
    }else{
        //只留最后一个
        stack.splice(0, endStackIndex); 
    }
    return result;
}
// @lc code=end

const result = trap([0,7,1,4,6]);
//const result = trap([5,2,1,2,1,5]);
//const result = trap([0,1,0,2,1,0,1,3,2,1,2,1]);

//const result = trap([0,1,0,2]);
console.log("result", result);