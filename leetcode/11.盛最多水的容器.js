/*
 * @lc app=leetcode.cn id=11 lang=javascript
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
    let left = 0;
    let right = height.length - 1;
    let maxArea = 0;
    while( right > left ){
        let area = 0;
        let dis = right - left;
        if( height[left] > height[right] ){
            area = dis * height[right];
            right--;
        }else{
            area = dis * height[left];
            left++;
        }
        if( area > maxArea ){
            maxArea = area;
        }
    }
    return maxArea;
};

//暴力法
var maxArea1 = function(height) {
    let maxArea = 0;
    for (let i = 0; i < height.length-1; i++) {
        for (let j = i; j < height.length; j++) {
            let area = Math.min( height[i], height[j] ) * (j - i);
            if( area > maxArea ){
                maxArea = area;
            }
        }
    }
    return maxArea;
};
// @lc code=end

