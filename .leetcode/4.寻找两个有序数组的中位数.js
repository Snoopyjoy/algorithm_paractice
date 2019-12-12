/*
 * @lc app=leetcode.cn id=4 lang=javascript
 *
 * [4] 寻找两个有序数组的中位数
 */

// @lc code=start
/**
 * @param {number[]} A
 * @param {number[]} B
 * @return {number}
 */
var findMedianSortedArrays = function(A, B) {
    let m = A.length;
    let n = B.length;
    if(m>n){    //长度小的数组为A
        return findMedianSortedArrays(B, A);
    }
    let maxLeft1, minRight1, maxLeft2, minRight2,
    c1,c2,lo = 0, hi = 2*m;
    while( lo <= hi ){
        //二分法求得第一个数组的切分位置
        c1 = ((lo + hi)/2)>>0;
        c2 = m + n - c1;
        maxLeft1 = c1 === 0 ? -Number.MAX_VALUE : A[((c1  - 1)/2)>>0];
        minRight1 = c1 === 2*m ? Number.MAX_VALUE : A[(c1/2)>>0];
        maxLeft2 = c2 === 0 ? -Number.MAX_VALUE : B[((c2  - 1)/2)>>0];
        minRight2 = c2 === 2*n ? Number.MAX_VALUE : B[(c2/2)>>0];

        // 目标 maxLeft1 < minRight1, maxLeft2 < minRight2
        if(maxLeft2 > minRight1 ){ //切分位置需要左移动
            lo = c1 + 1;
        }else if(  maxLeft1 > minRight2 ){//切分位置需要向右移动
            hi = c1 - 1;
        }else{
            break;
        }
    }
    Array.isArray
    return (Math.max(maxLeft2,maxLeft1) + 
    Math.min(minRight1, minRight2))/2;
};
// @lc code=end

//console.log( findMedianSortedArrays([3],[-2,-1]) )