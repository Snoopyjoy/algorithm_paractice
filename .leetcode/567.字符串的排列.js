/*
 * @lc app=leetcode.cn id=567 lang=javascript
 *
 * [567] 字符串的排列
 */

// @lc code=start
/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
var checkInclusion = function(s1, s2) {
    let s1Count = getEmptyCharArr();
    for (let i = 0; i < s1.length; i++) {
        let code = s1.charCodeAt(i) - 97;
        s1Count[code]++;
    }
    let s2TempCount = getEmptyCharArr();
    let s1Len = s1.length;
    for( let j = 0; j < s2.length; j++ ){
        let code = s2.charCodeAt(j) - 97;
        s2TempCount[code]++;
        if( j >= s1Len ){
            // 已经滑过的字符
            let passedCode = s2.charCodeAt(j - s1Len) - 97;
            s2TempCount[passedCode]--;
        }
        if( isEquelArray(s1Count, s2TempCount) ) return true;
    }
    return false;
};

function getEmptyCharArr(){
    return Array.from( Array(26) ).map(()=>0);
}

function isEquelArray( arrA, arrB ){
    return arrA.every( (v,i)=> v===arrB[i] );
}
// @lc code=end

//checkInclusion("ab","eidbaooo")

