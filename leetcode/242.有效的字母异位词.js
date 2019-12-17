/*
 * @lc app=leetcode.cn id=242 lang=javascript
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
    if( s.length !== t.length ) return false;
    let sCount = {};
    let tCount = {};
    countChar( s, sCount );
    countChar( t, tCount );
    const charsS = Object.keys(sCount);
    const charst = Object.keys(tCount);
    if( charsS.length !== charst.length ) return false;
    let result = charsS.every( (sKey)=>{
        return tCount[sKey] === sCount[sKey];
    } );
    return result;
};

function countChar( str, holder ){
    for (let index = 0; index < str.length; index++) {
        const char = str[index];
        if( holder[char] ){
            holder[char]++;
        }else{
            holder[char] = 1;
        }
    }
}
// @lc code=end

