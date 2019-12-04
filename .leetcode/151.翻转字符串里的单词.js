/*
 * @lc app=leetcode.cn id=151 lang=javascript
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function(s) {
    let tempS = s.trim();
    tempS = tempS.replace(/ +/g," ");
    return tempS.split(' ').reverse().join(' ');
};
// @lc code=end

