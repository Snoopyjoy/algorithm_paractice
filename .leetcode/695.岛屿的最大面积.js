/*
 * @lc app=leetcode.cn id=695 lang=javascript
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
/**
 * @param {number[][]} grid
 * @return {number}
 */
var maxAreaOfIsland = function(grid) {
    let memo = [];
    let maxArea = 0;
    grid.forEach(( line, i ) => {
        line.forEach( (val, j) => {
            let sumObj = {
                area: 0
            };
            if ( ( !memo[i] || !memo[i][j] ) && val === 1) {
                countArea ( grid, memo, i, j, sumObj )
            }
            if( sumObj.area > maxArea ) maxArea = sumObj.area;
        });
    });
    return maxArea;
};

function countArea ( grid, memo, i, j, sumObj ){
    if ( !memo[i] ) memo[i] = [];
    memo[i][j] = true;
    if ( grid[i][j] === 1 ) {
        sumObj.area += 1;
        // 上
        if ( i > 0 ) {
            if ( !memo[i-1] || !memo[i-1][j] ) {
                countArea ( grid, memo, i - 1, j, sumObj )
            }
        }
        // 下
        if ( i < grid.length - 1 ){
            if ( !memo[i+1] || !memo[i+1][j] ) {
                countArea ( grid, memo, i + 1, j, sumObj )
            }
        }
        // 左
        if ( j > 0 ){
            if ( !memo[i][j - 1] ) {
                countArea ( grid, memo, i, j - 1, sumObj )
            }
        }
        // 右
        if ( j < grid[i].length - 1 ){
            if ( !memo[i][j + 1] ) {
                countArea ( grid, memo, i, j + 1, sumObj )
            }
        }
    }
    return;
}
// @lc code=end

