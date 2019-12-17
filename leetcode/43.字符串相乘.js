/*
 * @lc app=leetcode.cn id=43 lang=javascript
 *
 * [43] 字符串相乘
 */

// @lc code=start
/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var multiply = function(num1, num2) {
    let resultArr = [];
    for (let i = num1.length - 1; i >= 0; i--) {
        for (let j = num2.length - 1; j >= 0; j--) {
            let n1 = Number(num1[i]);
            let n2 = Number(num2[j]);
            let pos = (num1.length - i) + (num2.length - j) - 2;
            let posV = resultArr[pos] || 0;
            let tempRst = n1 * n2;
            tempRst += posV;
            const lNum = tempRst % 10;
            const hNum = (tempRst / 10)>>0;
            resultArr[pos] = lNum;
            if( hNum > 0 ){
                resultArr[pos+1] = resultArr[pos+1] || 0;
                resultArr[pos+1] += hNum;
            }
        }
    }
    let temp = "";
    let valid = false;
    for (let i = resultArr.length - 1; i >= 0; i--) {
        let v = resultArr[i] || 0;
        if( !valid && v > 0   ){
           valid = true;
        }
        if (valid) temp += v;
    }
    temp = temp === "" ? "0" : temp;
    return temp;
};
// @lc code=end

