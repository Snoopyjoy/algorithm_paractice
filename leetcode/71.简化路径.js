/*
 * @lc app=leetcode.cn id=71 lang=javascript
 *
 * [71] 简化路径
 */

// @lc code=start
/**
 * @param {string} path
 * @return {string}
 */
var simplifyPath = function(path) {
    //path = path.replace(/\/+/g,"/");
    let pathArr = path.split('/');
    let i = pathArr.length - 1;
    let result = '';
    let backLevel = 0;
    while (i >= 0) {
        let curDir = pathArr[i];
        if( curDir !== '' ){
            if ( curDir === '..' ) {
                backLevel++;
            } else if ( curDir === '.') {
                // do nothing;
            } else {
                if ( backLevel > 0 ) {
                    backLevel--;
                } else {
                    result = '/' + curDir + result;
                }
            }
        }
        i--;
    }
    if( result === '' ) result = "/";
    return result;
};
// @lc code=end

