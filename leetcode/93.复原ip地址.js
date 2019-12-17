/*
 * @lc app=leetcode.cn id=93 lang=javascript
 *
 * [93] 复原IP地址
 */

// @lc code=start
/**
 * @param {string} s
 * @return {string[]}
 */
const MAX_VAL = 255;
const MIN_S = 1;
const MAX_S = 3;

var restoreIpAddresses = function(s) {
    let result = new Set();
    splitAddress(s, 4, '', result);
    return [...result];
};

function splitAddress( s, steps, temp, result ) {
    if( steps === 0 && s === '' ){
        result.add(temp);
        return;
    }
    // 判断字符是否足够继续分割
    let avgS = s.length / steps;
    if( avgS > MAX_S || avgS < MIN_S ) return;
    let r1 = s.substring(0,1);
    if( isValidNum(r1) ){
        let newTemp = temp === "" ? r1 : temp + "." + r1;
        splitAddress(s.substring(1), steps-1, newTemp, result);
    }
    let r2 = s.substring(0,2);
    if( isValidNum(r2) ){
        let newTemp = temp === "" ? r2 : temp + "." + r2;
        splitAddress(s.substring(2), steps-1, newTemp, result);
    }
    let r3 = s.substring(0,3);
    if( isValidNum(r3) ){
        let newTemp = temp === "" ? r3 : temp + "." + r3;
        splitAddress(s.substring(3), steps-1, newTemp, result);
    }
}

function isValidNum( num ){
    let _num = Number(num);
    return (num.length >1 && num[0] !== "0" || num.length === 1) && _num >= 0 && _num <= MAX_VAL;
}
// @lc code=end

