/*
 * @lc app=leetcode.cn id=121 lang=javascript
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let minPrice = Number.MAX_VALUE;
    let benifit = 0;
    prices.forEach( (price) => {
        if (price < minPrice) {
            minPrice = price;
        }else if (price > minPrice) {
            let tempBenifit = price - minPrice;
            if (tempBenifit > benifit) {
                benifit = tempBenifit;
            }
        }
    } );
    return benifit;
};
// @lc code=end

