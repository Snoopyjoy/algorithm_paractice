/*
 * @lc app=leetcode.cn id=122 lang=javascript
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let benifit = 0;
    if( prices.length === 0 ) return 0;
    prices.reduce( (pre, next)=>{
        if( next > pre ){
            benifit += next - pre;
        }
        return next;
    })
    return benifit;
};
// @lc code=end

