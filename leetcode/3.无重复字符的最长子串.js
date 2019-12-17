/*
 * @lc app=leetcode.cn id=3 lang=javascript
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    //滑动窗口法
    let charMap = new Map();
    let maxLen = 0;
    /* 窗口起始位置 */
    let startIndex = 0;
    let len = 0;
    let i = 0;
    while (i < s.length) {
        let char = s[i];
        //窗口中已存在该字符
        if (charMap.has(char) && charMap.get(char) >= startIndex) {
            startIndex = charMap.get(char) + 1;
            len = i - charMap.get(char);
        }
        else {
            len++;
        }
        if( len > maxLen ) maxLen = len;
        charMap.set(char, i);
        i++;
    }
    return maxLen;
};
// @lc code=end

