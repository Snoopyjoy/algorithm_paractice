/*
 * @lc app=leetcode.cn id=20 lang=javascript
 *
 * [20] 有效的括号
 */

// @lc code=start
/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    let len = s.length;
    if( len % 2 === 1 ) return false;
    let stack = [];
    for (let i = 0; i < len; i++) {
        const char = s[i];
        if( isLeft(char) ){
            stack.push(char);
        }else{
            let top = stack.pop();
            if( !isMatch(top, char) ) return false;
        }
    }
    return stack.length === 0;
};
function isLeft( s ){
    return s === "{" || s === "(" || s === "[" 
}

function isMatch( a, b ){
    return a === "(" && b === ")" ||
        a === "{" && b === "}" ||
        a === "[" && b === "]"
}
// @lc code=end

