/*
 * @lc app=leetcode.cn id=10 lang=javascript
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
/**
 * @param {string} s
 * @param {string} p
 * @return {boolean}
 */
var isMatch = function(s, p) {
    const memo = [];
    return dp( memo, 0, 0, s, p );
};

function dp( memo, i , j, text, pattern ){
    if( memo[i] !== undefined && memo[i][j] !== undefined ){
        return memo[i][j];
    }
    if( memo[i] === undefined ) memo[i] = [];
    let ans = false;
    if( j === pattern.length ){
        ans = i === text.length;
    }else{
        let firstMatch = i < text.length && (text[i] === pattern[j] ||
            pattern[j] === ".");
        if( j+1 < pattern.length && pattern[j+1] === "*" ){
            ans =  dp( memo, i, j+2, text, pattern ) || //*代表0个前字符
            ( firstMatch && dp( memo, i+1, j , text, pattern) );
        }else{
            ans = firstMatch && dp(memo, i+1, j+1, text, pattern);
        }
    }
    memo[i][j] = ans;
    return ans;
}
// @lc code=end

