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
        let temp = m;
        m = n;
        n = temp;
        let temp1 = A;
        A = B;
        B = temp1;
    }
    let iMin = 0, iMax = m;
    // 两个数组的中间长度（奇数为中间位置+1）（偶数为偏右的位置）
    let halfLen = ((m+n+1)/2)>>0;
    //双指针
    while(iMin <= iMax){
        //i 为A数组的中间位置（偶数长度为偏右的位置）
        let i = ((iMin + iMax)/2)>>0;
        //j为拼接后的数组的中间位置和A数组中心位置的偏移量
        let j = halfLen - i;
        if( i < iMax && B[j-1] > A[i] ){
           iMin = i + 1;
        }
        else if( i > iMin && A[i-1] > B[j] ){
            iMax = i -1;
        }
        else{
            let maxLeft = 0;
            if(i==0){ maxLeft = B[j-1]; }
            else if(j==0){ maxLeft = A[i-1]; }
            else{ maxLeft = Math.max(A[i-1], B[j-1]); }
            if( (m+n)%2 === 1 ) return maxLeft;
            let minRight = 0;
            if(i == m){ minRight = B[j]; }
            else if( j == n ){ minRight = A[i]; }
            else{ minRight = Math.min(B[j], A[i]); }
            
            return ( maxLeft + minRight ) / 2;
        }
    }
    return 0;
};
// @lc code=end

