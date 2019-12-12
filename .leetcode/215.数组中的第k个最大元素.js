/*
 * @lc app=leetcode.cn id=215 lang=javascript
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */

var findKthLargest = function(nums, k) {
    let startIndex = 0;
    let endIndex = nums.length - 1;
    let target = nums.length - k;
    while (startIndex <= endIndex) {
        let pivotIndex = partition(nums, startIndex, endIndex);
        if ( pivotIndex > target ) {
            endIndex = pivotIndex - 1;
        } else if( pivotIndex < target ) {
            startIndex = pivotIndex + 1;
        } else {
            return nums[pivotIndex];
        }
    }
    return nums[0];
};

// 分割
function partition(nums, startIndex, endIndex) {
    // 哨兵
    const pivot = nums[startIndex];
    
    let left = startIndex;
    let right = endIndex;
    while (left < right){
        while (left < right && nums[right] > pivot ) right --;
        while (left < right && nums[left] <= pivot ) left ++;
        if (left < right) {
            let temp = nums[left];
            nums[left] = nums[right];
            nums[right] = temp;
        }
    }
    nums[startIndex] = nums[left];
    nums[left] = pivot;
    return left;
}

//堆排序
function findKthLargest1(nums, k) {
    //构建最大堆
}

function ajustDown( nums, parentIndex, len ) {
    let childIndex = 2 * parentIndex + 1;
    let parentNodeVal = nums[parentIndex];
    while (childIndex < len) {
        let maxChildVal = nums[childIndex];
        if( childIndex + 1 < len && maxChildVal > nums[childIndex + 1] ) {
            childIndex++;
            maxChildVal = nums[childIndex];
        }
        // 父节点大于所有子节点
        if( parentNodeVal > maxChildVal ) break;
        // 父节点与子节点交换位置

    }
}
// @lc code=end

let result = findKthLargest([3,2,3,1,2,4,5,5,6],4 );
console.log(result);

