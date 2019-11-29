/*
 * @lc app=leetcode.cn id=5 lang=javascript
 *
 * [5] 最长回文子串
 */

// @lc code=start
/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
    return dp(s);
};

//双指针 中心扩散
function dobulePoint(s){
    if(!s || s.length < 2) return s;
    let leftPoint = 0;
    let rightPoint = 1;
    let subStrPos = [0,0];
    while( rightPoint < s.length ){
        if(s[leftPoint] === s[rightPoint]){
            expand( s, leftPoint, rightPoint, subStrPos );
        }
        if( rightPoint - leftPoint > 1 ){
            leftPoint++;
        }else{
            rightPoint++;
        }
    }
    return s.substring( subStrPos[0], subStrPos[1] + 1 );
}

function expand( s, leftPoint, rightPoint, subStrPos ){
    while( leftPoint > -1 && rightPoint < s.length && s[leftPoint] === s[rightPoint] ){
        let curMax = subStrPos[1] - subStrPos[0] + 1;
        if( rightPoint - leftPoint + 1 > curMax ){
            subStrPos[0] = leftPoint;
            subStrPos[1] = rightPoint;
        }
        leftPoint--;
        rightPoint++;
    }

}

//动态规划
function dp(s){
    if( s.length <= 1 ) return s;
    let memo = [];
    let maxLen = 1;
    let result = s[0];

    for (let right = 1; right < s.length; right++) {  //右边界
        for( let left = 0; left < right; left++ ){ //起始位置
            if( memo[left] === undefined ) memo[left] = [];
            if( memo[left+1] === undefined ) memo[left+1] = [];
            if( s[right] === s[left] && ( right - left <= 2 || memo[left+1][right-1] )){
                memo[left][right] = true;
                if( right - left + 1 > maxLen ){
                    maxLen = right - left + 1;
                    result = s.substring(left,right+1);
                }
            }
        }
    }
    return result;
}
// @lc code=end

